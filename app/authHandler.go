package app

import (
	"bankingAuth/domain"
	"bankingAuth/errs"
	"bankingAuth/service"
	"encoding/json"
	"log"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func (handler AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest domain.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Print("Failed to parse the request body")
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequest("invalid request").AsMessage())
		return
	}
	loginResponse, loginError := handler.service.Login(loginRequest)
	if loginError != nil {
		writeResponse(w, http.StatusUnauthorized, loginError.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, loginResponse)
}

func writeResponse(w http.ResponseWriter, code int, t interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(t)
	if err != nil {
		panic(err)
	}
}
