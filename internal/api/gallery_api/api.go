package gallery_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	engine *gin.Engine
}

func (api *API) Run() error {
	api.engine = gin.Default()
	api.SetupGeneralRouter()
	return api.engine.Run()
}

func (api *API) SetupGeneralRouter() {
	api.engine.GET("/ping", Ping)
	api.engine.GET("/healthcheck", HealthCheck)
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "PONG")
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
