package controllers

import (
	"net/http"
	"rest-api/request"
	"rest-api/response"
	"rest-api/services"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(*gin.Context)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(_s services.AuthService) AuthController {
	return authController{
		authService: _s,
	}
}

func (_c authController) Login(c *gin.Context) {
	var req request.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Field username or password can't be empty",
		})
		c.Abort()
		return
	}

	result, err := _c.authService.Login(req)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Username and Password do not match",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		c.Abort()
		return
	}

	isMatch := utils.CheckHash(req.Password, result.Password)
	if !isMatch {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username and Password do not match",
		})
		c.Abort()
		return
	}

	token, err := utils.GenerateJWT(result)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": "Error generate token",
		})
		c.Abort()
		return
	}

	var res response.LoginResponse
	res = res.LoginTransform(result, token)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
}
