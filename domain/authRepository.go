package domain

import "bankingAuth/errs"

type AuthRepository interface {
	FindBy(username string, password string) (*Login, *errs.AppError)
}
