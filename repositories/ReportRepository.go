package repositories

import (
	"rest-api/response"
	"rest-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type reportRepository struct {
	DB *gorm.DB
}

type ReportRepository interface {
	ReportMerchant(*gin.Context) (paginate.Page, error)
	ReportOutlet(*gin.Context) (paginate.Page, error)
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return reportRepository{
		DB: db,
	}
}

func (_r reportRepository) ReportMerchant(c *gin.Context) (paginate.Page, error) {
	var err error
	var size int
	var page int
	query := c.Request.URL.Query()

	month := query.Get("month")
	from, to := utils.GetFirstAndLastDayofMonth(month)

	periodes := utils.GetPeriode(month)
	page, _ = strconv.Atoi(query.Get("page"))
	size, _ = strconv.Atoi(query.Get("size"))
	if page > 0 {
		page = page * size
	}

	userID, _ := c.Get("user_id")
	model := _r.DB.Raw(`
		SELECT
			m.merchant_name,
			SUM(t.bill_total) AS omzet,
			SUBSTRING(t.created_at,1,10) AS date
		FROM transactions t
		JOIN merchants m ON m.id = t.merchant_id
		WHERE m.user_id = ? AND t.created_at BETWEEN ? AND ?
		GROUP BY date
	`, userID, from, to)
	pg := paginate.New()
	pgResult := pg.Response(model, c.Request, &[]response.ReportMerchantResponse{})

	var result []response.ReportMerchantResponse
	model.Scan(&result)

	var datas []response.ReportMerchantResponse
	var dt []string
	for i := page; i < (page + size); i++ {
		for _, v := range result {
			if periodes[i] == v.Date {
				datas = append(datas, v)
				dt = append(dt, v.Date)
			}
		}
	}

	for i := page; i < (page + size); i++ {
		if !utils.StringInSlice(periodes[i], dt) {
			datas = append(datas, response.ReportMerchantResponse{
				Date:         periodes[i],
				MerchantName: result[0].MerchantName,
				Omzet:        0,
			})
		}
	}

	totalPage := len(periodes) / size
	maxPage := totalPage
	if totalPage > 1 {
		maxPage = totalPage - 1
	}
	pgResult.Items = datas
	pgResult.MaxPage = int64(maxPage)
	pgResult.TotalPages = int64(totalPage)
	pgResult.Total = int64(len(periodes))
	return pgResult, err
}

func (_r reportRepository) ReportOutlet(c *gin.Context) (paginate.Page, error) {
	var err error
	var size int
	var page int
	query := c.Request.URL.Query()

	month := query.Get("month")
	from, to := utils.GetFirstAndLastDayofMonth(month)

	periodes := utils.GetPeriode(month)
	page, _ = strconv.Atoi(query.Get("page"))
	size, _ = strconv.Atoi(query.Get("size"))
	if page > 0 {
		page = page * size
	}

	userID, _ := c.Get("user_id")
	model := _r.DB.Raw(`
		SELECT
			m.merchant_name,
			o.outlet_name,
			SUM(t.bill_total) AS omzet,
			SUBSTRING(t.created_at,1,10) AS date
		FROM transactions t
		JOIN merchants m ON m.id = t.merchant_id
		JOIN outlets o ON o.id = t.outlet_id
		WHERE m.user_id = ? AND t.created_at BETWEEN ? AND ?
		GROUP BY date
	`, userID, from, to)
	pg := paginate.New()
	pgResult := pg.Response(model, c.Request, &[]response.ReportOutletResponse{})

	var result []response.ReportOutletResponse
	model.Scan(&result)

	var datas []response.ReportOutletResponse
	var dt []string
	for i := page; i < (page + size); i++ {
		for _, v := range result {
			if periodes[i] == v.Date {
				datas = append(datas, v)
				dt = append(dt, v.Date)
			}
		}
	}

	for i := page; i < (page + size); i++ {
		if !utils.StringInSlice(periodes[i], dt) {
			datas = append(datas, response.ReportOutletResponse{
				Date:         periodes[i],
				MerchantName: result[0].MerchantName,
				OutletName:   result[0].OutletName,
				Omzet:        0,
			})
		}
	}

	totalPage := len(periodes) / size
	maxPage := totalPage
	if totalPage > 1 {
		maxPage = totalPage - 1
	}
	pgResult.Items = datas
	pgResult.MaxPage = int64(maxPage)
	pgResult.TotalPages = int64(totalPage)
	pgResult.Total = int64(len(periodes))
	return pgResult, err
}
