package services

import (
	"rest-api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

type ReportService interface {
	ReportMerchant(*gin.Context) (paginate.Page, error)
	ReportOutlet(*gin.Context) (paginate.Page, error)
}

type reportService struct {
	reportRepository repositories.ReportRepository
}

func NewReportService(_s repositories.ReportRepository) ReportService {
	return reportService{
		reportRepository: _s,
	}
}

func (_s reportService) ReportMerchant(c *gin.Context) (paginate.Page, error) {
	return _s.reportRepository.ReportMerchant(c)
}

func (_s reportService) ReportOutlet(c *gin.Context) (paginate.Page, error) {
	return _s.reportRepository.ReportOutlet(c)
}
