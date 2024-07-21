package auth_service

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
