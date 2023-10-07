package domain

type Claims struct {
	CustomerID string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Role       string   `json:"role"`
	Username   string   `json:"username"`
	Expiry     int64    `json:"exp"`
}

func (c Claims) IsUserRole() bool {
	return c.Role == "user"
}

func (c Claims) isValidAccountID(accountID string) bool {
	if accountID != "" {
		for _, x := range c.Accounts {
			if x == accountID {
				return true
			}
		}
		return false
	}
	return true
}

func (c Claims) IsRequestVerifiedWithTokenClaims(params map[string]string) bool {
	if c.CustomerID != params["customer_id"] {
		return false
	}
	if !c.isValidAccountID(params["account_id"]) {
		return false
	}
	return true
}
