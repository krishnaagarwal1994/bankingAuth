package domain

type LoginResponse struct {
	Token string `json:"auth_token"`
}
