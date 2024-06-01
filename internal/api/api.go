package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/api/gallery_api"
	"go-mercury/internal/data/gallery_db"
	"go-mercury/internal/service/gallery_service"
)

type App struct {
	engine         *gin.Engine
	galleryHandler *gallery_api.GalleryHandler
}

func NewApp() App {
	return App{}
}

func (api *App) Run() error {
	api.engine = gin.Default()
	err := api.SetupDependencies()
	if err != nil {
		log.Errorf("err api.SetupDependencies: %v", err)
		return err
	}
	api.SetupRouter()
	return api.engine.Run()
}

func (api *App) SetupDependencies() error {
	galleryDB := gallery_db.NewDB("127.0.0.1", 5432, "username", "password", "gallery_db")
	galleryService := gallery_service.NewService(galleryDB)
	api.galleryHandler = gallery_api.NewGalleryHandler(&galleryService)

	return nil
}

func (api *App) SetupRouter() {
	api.engine.GET("/ping", Ping)
	api.engine.GET("/healthcheck", HealthCheck)

	api.engine.GET("/products/:id", api.galleryHandler.GetProduct)
	api.engine.DELETE("/products/:id", api.galleryHandler.DeleteProduct)
	api.engine.POST("/products", api.galleryHandler.CreateProduct)
	api.engine.PATCH("/products", api.galleryHandler.UpdateProduct)
}
