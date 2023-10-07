package service

import (
	"bankingAuth/domain"
	"bankingAuth/errs"
	"encoding/json"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type DefaultAuthService struct {
	repository      domain.AuthRepository
	rolePermissions domain.RolePermissions
}

func (service DefaultAuthService) Login(loginRequest domain.LoginRequest) (*domain.LoginResponse, *errs.AppError) {
	login, err := service.repository.FindBy(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return nil, err
	}
	token, tokenError := login.GenerateToken()
	if tokenError != nil {
		return nil, err
	}
	loginResponse := domain.LoginResponse{Token: *token}
	return &loginResponse, nil
}

func (service DefaultAuthService) Verify(urlParams map[string]string) (bool, *errs.AppError) {
	// Convert the string token to jwt struct
	var jwtToken *jwt.Token
	var jwtTokenError *errs.AppError
	if jwtToken, jwtTokenError = service.jwtTokenFromString(urlParams["token"]); jwtTokenError != nil {
		return false, jwtTokenError
	}
	// Checking the validity of the token, this verifies the expiry time and the signature of the token
	if jwtToken.Valid {
		// type cast the token claims to jwt.MapClaims
		mapClaims := jwtToken.Claims.(jwt.MapClaims)
		// Converting the token claims to Claims struct
		claims, err := service.buildClaimsFromJWTMapClaims(mapClaims)
		if err != nil {
			return false, err
		}
		if claims.IsUserRole() && !claims.IsRequestVerifiedWithTokenClaims(urlParams) {
			return false, errs.NewForbiddenError("unauthorized")
		}
		isAuthorized := service.rolePermissions.IsAuthorizedFor(claims.Role, urlParams["routeName"])
		return isAuthorized, nil
	} else {
		return false, errs.NewBadRequest("invalid token")
	}
}

func (service DefaultAuthService) buildClaimsFromJWTMapClaims(claims jwt.MapClaims) (*domain.Claims, *errs.AppError) {
	bytes, err := json.Marshal(claims)
	if err != nil {
		log.Print("failed to marshel the json")
		return nil, errs.NewAppError("failed to marshel the json")
	}
	var myClaims *domain.Claims
	if err := json.Unmarshal(bytes, &myClaims); err != nil {
		log.Print("failed to unmarshel the json with error" + err.Error())
		return nil, errs.NewAppError("failed to unmarshel the json")
	}
	return myClaims, nil
}

func (service DefaultAuthService) jwtTokenFromString(tokenString string) (*jwt.Token, *errs.AppError) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SECRET), nil
	})
	if err != nil {
		log.Print("Error while parsing token")
		return nil, errs.NewAppError("Error while parsing JWT Token")
	}
	return token, nil
}

func NewAuthService(repository domain.AuthRepository, rolePermissions domain.RolePermissions) AuthService {
	return DefaultAuthService{repository: repository, rolePermissions: rolePermissions}
}
