package auth_service

import (
	"github.com/dgrijalva/jwt-go"

	"go-mercury/internal/data/gallery_db"
)

type JWTClaims struct {
	Username string                `json:"username"`
	Roles    []gallery_db.UserRole `json:"roles"`
	jwt.StandardClaims
}
