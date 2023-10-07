package domain

import "bankingAuth/errs"

type LoginRequest struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
}

func (request LoginRequest) Validate() *errs.AppError {
	if request.Username == "" || request.Password == "" {
		return errs.NewBadRequest("invalid request")
	} else {
		return nil
	}
}
