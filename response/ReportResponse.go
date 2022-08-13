package response

type ReportMerchantResponse struct {
	Date         string `json:"date"`
	MerchantName string `json:"merchant_name"`
	Omzet        int64  `json:"omzet"`
}

type ReportOutletResponse struct {
	Date         string `json:"date"`
	MerchantName string `json:"merchant_name"`
	OutletName   string `json:"outlet_name"`
	Omzet        int64  `json:"omzet"`
}
