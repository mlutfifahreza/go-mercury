package auth_api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/api/general"
	"go-mercury/internal/service/auth_service"
	"go-mercury/pkg/constant"
	"go-mercury/pkg/util"
)

type AuthHandler struct {
	authService *auth_service.Service
}

func NewAuthHandler(authService *auth_service.Service) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	reqBody, err := util.ParseRequestBody[Credentials](c)
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, err)
		return
	}

	jwtString, err := h.authService.UserRegister(reqBody.Username, reqBody.Password)
	if err != nil {
		log.WithError(err).Error("authService.UserLogin")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	setResponseJWT(c, jwtString)
	general.CreateSuccessResponse(c, nil)

}

func (h *AuthHandler) Login(c *gin.Context) {
	reqBody, err := util.ParseRequestBody[Credentials](c)
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, err)
		return
	}

	jwtString, err := h.authService.UserLogin(reqBody.Username, reqBody.Password)
	if err != nil {
		if errors.Is(err, constant.UserNotFoundError) || errors.Is(err, constant.WrongCredentialsError) {
			general.CreateFailResponse(c, http.StatusUnauthorized, err)
			return
		}

		log.WithError(err).Error("authService.UserLogin")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	setResponseJWT(c, jwtString)
	general.CreateSuccessResponse(c, nil)
}

func (h *AuthHandler) UserData(c *gin.Context) {
	tokenCookie, err := c.Cookie(CookieKeyJWT)
	if err != nil {
		general.CreateFailResponse(c, http.StatusUnauthorized, err)
		return
	}

	claim, err := h.authService.CheckJWT(tokenCookie)
	if err != nil {
		general.CreateFailResponse(c, http.StatusUnauthorized, err)
		return
	}

	userTab, err := h.authService.GetUser(claim.Username)
	if err != nil {
		log.WithError(err).Error("h.authService.GetUser(claim.Username)")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	userData := UserData{Username: userTab.Username}
	userData.SetRoles(userTab.Roles)
	general.CreateSuccessResponse(c, userData)
}

func setResponseJWT(c *gin.Context, jwtString string) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    jwtString,
		Path:     "/",
		MaxAge:   86400, // 1 day
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(c.Writer, &cookie)
}
