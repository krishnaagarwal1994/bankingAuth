package domain

import (
	"bankingAuth/errs"
	"database/sql"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TOKEN_DURATION = time.Hour
const HMAC_SECRET = "sdfvgbhgnfsadfghdsfghjfddfg"

type Login struct {
	Username   string         `json:"username"`
	CustomerId sql.NullString `json:"customer_id"`
	Accounts   sql.NullString `json:"account_numbers"`
	Role       string         `json:"role"`
}

func (l Login) GenerateToken() (*string, *errs.AppError) {
	var claims jwt.MapClaims
	if l.Accounts.Valid && l.CustomerId.Valid {
		claims = l.claimsForUser()
	} else {
		claims = l.claimsForAdmin()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(HMAC_SECRET))
	if err != nil {
		return nil, errs.NewInternalServerError("Token generation failed")
	}
	return &tokenString, nil
}

func (l Login) claimsForUser() jwt.MapClaims {
	accounts := strings.Split(l.Accounts.String, ",")
	claims := jwt.MapClaims{
		"customer_id": l.CustomerId.String,
		"role":        l.Role,
		"username":    l.Username,
		"accounts":    accounts,
		"exp":         time.Now().Add(TOKEN_DURATION).Unix(),
	}
	return claims
}

func (l Login) claimsForAdmin() jwt.MapClaims {
	claims := jwt.MapClaims{
		"role":     l.Role,
		"username": l.Username,
		"exp":      time.Now().Add(TOKEN_DURATION).Unix(),
	}
	return claims
}
