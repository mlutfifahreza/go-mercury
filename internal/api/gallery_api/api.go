package gallery_api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/data/gallery_db"
	"go-mercury/internal/service/gallery_service"
)

type API struct {
	engine         *gin.Engine
	galleryHandler *GalleryHandler
}

func (api *API) Run() error {
	api.engine = gin.Default()
	err := api.SetupDependencies()
	if err != nil {
		log.Errorf("err api.SetupDependencies: %v", err)
		return err
	}
	api.SetupRouter()
	return api.engine.Run()
}

func (api *API) SetupDependencies() error {
	galleryDB := gallery_db.NewDB("127.0.0.1", 5432, "username", "password", "gallery_db")
	galleryService := gallery_service.NewService(galleryDB)
	api.galleryHandler = NewGalleryHandler(&galleryService)

	return nil
}

func (api *API) SetupRouter() {
	api.engine.GET("/ping", Ping)
	api.engine.GET("/healthcheck", HealthCheck)

	api.engine.GET("/products/:id", api.galleryHandler.getPoll)
	api.engine.POST("/products/", api.galleryHandler.createPoll)
}
