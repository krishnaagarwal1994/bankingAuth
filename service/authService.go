package service

import (
	"bankingAuth/domain"
	"bankingAuth/errs"
)

type AuthService interface {
	Login(loginRequest domain.LoginRequest) (*domain.LoginResponse, *errs.AppError)
	Verify(urlParams map[string]string) (bool, *errs.AppError)
}
