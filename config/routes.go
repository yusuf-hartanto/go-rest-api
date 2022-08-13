package config

import (
	"rest-api/controllers"
	"rest-api/middleware"
	"rest-api/repositories"
	"rest-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) {
	router := gin.Default()

	// auth
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)
	public := router.Group("api")
	public.POST("/login", authController.Login)

	// report
	reportRepository := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepository)
	reportController := controllers.NewReportController(reportService)
	protected := router.Group("api/report")
	protected.Use(middleware.Authentication())
	protected.GET("/merchant", reportController.ReportMerchant)
	protected.GET("/outlet", reportController.ReportOutlet)

	router.Run()
}
