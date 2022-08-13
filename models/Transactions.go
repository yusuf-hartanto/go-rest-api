package models

type Transactions struct {
	ID         int64   `gorm:"column:id;primaryKey;autoIncrement:true"`
	MerchantID int64   `gorm:"column:merchant_id"`
	OutletID   string  `gorm:"column:outlet_id"`
	BillTotal  float64 `gorm:"column:bill_total"`
	CreatedBy  int64   `gorm:"column:created_by"`
	CreatedAt  string  `gorm:"column:created_at"`
	UpdatedBy  int64   `gorm:"column:updated_by"`
	UpdatedAt  string  `gorm:"column:updated_at"`
}
