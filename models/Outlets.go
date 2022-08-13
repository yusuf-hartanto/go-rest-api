package models

type Outlets struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true"`
	MerchantID  int64  `gorm:"column:merchant_id"`
	OutletsName string `gorm:"column:outlet_name"`
	CreatedBy   int64  `gorm:"column:created_by"`
	CreatedAt   string `gorm:"column:created_at"`
	UpdatedBy   int64  `gorm:"column:updated_by"`
	UpdatedAt   string `gorm:"column:updated_at"`
}
