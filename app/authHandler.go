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
	// Variable declaration to hold the request body
	var loginRequest domain.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Print("Failed to parse the request body")
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequest("invalid request").AsMessage())
		return
	}

	// Performing validation on the mapped request body
	validationError := loginRequest.Validate()
	if validationError != nil {
		writeResponse(w, validationError.Code, validationError.AsMessage())
		return
	}

	// Passing the validated LoginRequest to the AuthService to perform login operation.
	loginResponse, loginError := handler.service.Login(loginRequest)
	if loginError != nil {
		writeResponse(w, http.StatusUnauthorized, loginError.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, loginResponse)
}

func (handler AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	urlParams := make(map[string]string)

	for k, _ := range r.URL.Query() {
		urlParams[k] = r.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {
		isAuthorized, err := handler.service.Verify(urlParams)
		if err != nil {
			log.Print("Failed to verify the auth token")
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			if isAuthorized {
				writeResponse(w, http.StatusOK, handler.authorizedResponse())
			} else {
				writeResponse(w, http.StatusForbidden, handler.unauthorizedResponse())
			}
		}
	} else {
		writeResponse(w, http.StatusForbidden, errs.AppError{Message: "missing token"}.AsMessage())
	}
}

func (handler AuthHandler) authorizedResponse() domain.VerifyTokenResponse {
	return domain.VerifyTokenResponse{IsAuthorized: true}
}

func (handler AuthHandler) unauthorizedResponse() domain.VerifyTokenResponse {
	return domain.VerifyTokenResponse{IsAuthorized: false}
}

func writeResponse(w http.ResponseWriter, code int, t interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(t)
	if err != nil {
		panic(err)
	}
}
