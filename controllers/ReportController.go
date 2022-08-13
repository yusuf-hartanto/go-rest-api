package controllers

import (
	"net/http"
	"rest-api/services"

	"github.com/gin-gonic/gin"
)

type ReportController interface {
	ReportMerchant(*gin.Context)
	ReportOutlet(*gin.Context)
}

type reportController struct {
	reportService services.ReportService
}

func NewReportController(_s services.ReportService) ReportController {
	return reportController{
		reportService: _s,
	}
}

func (_c reportController) ReportMerchant(c *gin.Context) {
	result, err := _c.reportService.ReportMerchant(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result)
}

func (_c reportController) ReportOutlet(c *gin.Context) {
	result, err := _c.reportService.ReportOutlet(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result)
}
