package app

import (
	"bankingAuth/domain"
	"bankingAuth/service"
	"database/sql"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	dbClient := getSQLClient()
	authRepository := domain.NewAuthRepository(dbClient)
	authService := service.NewAuthService(authRepository)
	authHandler := AuthHandler{service: authService}

	router.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	http.ListenAndServe("localhost:8081", router)
}

func getSQLClient() *sql.DB {
	db, err := sql.Open("mysql", "root:Gn1d0c@123@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
