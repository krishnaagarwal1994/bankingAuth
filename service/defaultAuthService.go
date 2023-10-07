package service

import (
	"bankingAuth/domain"
	"bankingAuth/errs"
)

type DefaultAuthService struct {
	repository domain.AuthRepository
}

func (service DefaultAuthService) Login(loginRequest domain.LoginRequest) (*domain.LoginResponse, *errs.AppError) {
	login, err := service.repository.FindBy(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return nil, err
	}
	token, tokenError := login.GenerateToken()
	if tokenError != nil {
		return nil, err
	}
	loginResponse := domain.LoginResponse{Token: *token}
	return &loginResponse, nil
}

// func (service DefaultAuthService) Verify(jwt string) (*error) {

// }

func NewAuthService(repository domain.AuthRepository) AuthService {
	return DefaultAuthService{repository: repository}
}
