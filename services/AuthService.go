package services

import (
	"rest-api/models"
	"rest-api/repositories"
	"rest-api/request"
)

type AuthService interface {
	Login(request.LoginReq) (models.User, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(_s repositories.AuthRepository) AuthService {
	return authService{
		authRepository: _s,
	}
}

func (_s authService) Login(req request.LoginReq) (models.User, error) {
	return _s.authRepository.Login(req)
}
