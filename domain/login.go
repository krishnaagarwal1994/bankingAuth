package domain

import "database/sql"

type Login struct {
	Username   string         `json:"username"`
	CustomerId sql.NullString `json:"customer_id"`
	Accounts   sql.NullString `json:"account_numbers"`
	Role       string         `json:"role"`
}

func (l Login) GenerateToken() (*string, *error) {
	token := "test token"
	return &token, nil
}
