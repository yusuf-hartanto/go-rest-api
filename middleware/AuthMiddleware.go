package middleware

import (
	"net/http"
	"rest-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is required",
			})
			c.Abort()
			return
		}
		splitBearer := strings.Split(authorization, " ")
		user, err := utils.ValidationToken(splitBearer[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("user_id", user.UserID)
		c.Next()
	}
}
