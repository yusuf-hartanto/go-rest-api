package models

type Merchants struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true"`
	UserID       int64  `gorm:"column:user_id"`
	MerchantName string `gorm:"column:merchant_name"`
	CreatedBy    int64  `gorm:"column:created_by"`
	CreatedAt    string `gorm:"column:created_at"`
	UpdatedBy    int64  `gorm:"column:updated_by"`
	UpdatedAt    string `gorm:"column:updated_at"`
}
