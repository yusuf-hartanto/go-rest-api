package models

type User struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true"`
	Name      string `gorm:"column:name"`
	Username  string `gorm:"column:user_name"`
	Password  string `gorm:"column:password"`
	CreatedBy int64  `gorm:"column:created_by"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedBy int64  `gorm:"column:updated_by"`
	UpdatedAt string `gorm:"column:updated_at"`
}
