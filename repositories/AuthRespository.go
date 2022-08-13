package repositories

import (
	"rest-api/models"
	"rest-api/request"

	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

type AuthRepository interface {
	Login(request.LoginReq) (models.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return authRepository{
		DB: db,
	}
}

func (_r authRepository) Login(req request.LoginReq) (models.User, error) {
	var err error

	var result models.User
	if err := _r.DB.Where("user_name = ?", req.Username).First(&result).Error; err != nil {
		return models.User{}, err
	}
	return result, err
}
