package domain

import (
	"bankingAuth/errs"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type AuthRepositoryDB struct {
	client *sql.DB
}

func (db AuthRepositoryDB) FindBy(username string, password string) (*Login, *errs.AppError) {
	sqlQuery := `select username, role, u.customer_id, group_concat(a.account_id) as account_numbers from users u LEFT JOIN accounts a ON a.customer_id = u.customer_id where username = ? and u.password = ? group by a.customer_id;`

	row := db.client.QueryRow(sqlQuery, username, password)
	var login Login
	err := row.Scan(&login.Username, &login.Role, &login.CustomerId, &login.Accounts)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewBadRequest("invalid credentials")
		} else {
			log.Print(err.Error())
			return nil, errs.NewDBError("unexpected db issue")
		}
	}
	return &login, nil
}

func NewAuthRepository(client *sql.DB) AuthRepository {
	return AuthRepositoryDB{client: client}
}
