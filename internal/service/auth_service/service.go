package auth_service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"go-mercury/internal/data/gallery_db"
	"go-mercury/pkg/constant"
)

type Service struct {
	db           *gallery_db.DB
	jwtKey       []byte
	passwordSalt string
}

func NewService(db *gallery_db.DB) Service {
	keyString := "ABCDE"
	return Service{
		db:     db,
		jwtKey: []byte(keyString),
	}
}

func (s Service) UserRegister(username string, password string) (string, error) {
	hashPassword, err := s.hashPassword(password)
	if err != nil {
		return "", err
	}

	err = s.db.CreateUserTab(gallery_db.User{
		Username:     username,
		PasswordHash: hashPassword,
		Roles:        nil,
	})
	if err != nil {
		return "", err
	}

	tokenString, err := s.generateJWT(username, []gallery_db.UserRole{})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s Service) UserLogin(username string, password string) (string, error) {
	userTab, err := s.db.GetUserTab(username)
	if err != nil {
		return "", err
	}

	if !s.checkPasswordHash(password, userTab.PasswordHash) {
		return "", constant.WrongCredentialsError
	}

	tokenString, err := s.generateJWT(userTab.Username, userTab.Roles)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s Service) CheckJWT(jwtString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, constant.WrongTokenError
	}

	return claims, nil

}

func (s Service) generateJWT(username string, roles []gallery_db.UserRole) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaims{
		Username: username,
		Roles:    roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		log.Errorf("token.SignedString: %v", err)
		return "", err
	}

	return tokenString, nil
}

func (s Service) hashPassword(password string) (string, error) {
	//saltedPassword := password + s.passwordSalt
	saltedPassword := password
	bytes, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), 14)
	return string(bytes), err
}

func (s Service) checkPasswordHash(password, hash string) bool {
	//saltedPassword := password + s.passwordSalt
	saltedPassword := password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedPassword))
	return err == nil
}
