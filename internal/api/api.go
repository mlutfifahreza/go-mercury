package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/api/product_api"
	"go-mercury/internal/api/store_api"
	"go-mercury/internal/data/gallery_db"
	"go-mercury/internal/service/gallery_service"
)

type App struct {
	engine         *gin.Engine
	productHandler *product_api.ProductHandler
	storeHandler   *store_api.StoreHandler
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
	api.productHandler = product_api.NewProductHandler(&galleryService)
	api.storeHandler = store_api.NewStoreHandler(&galleryService)

	return nil
}

func (api *App) SetupRouter() {
	api.engine.GET("/ping", Ping)
	api.engine.GET("/healthcheck", HealthCheck)

	api.engine.GET("/products/:id", api.productHandler.GetProduct)
	api.engine.DELETE("/products/:id", api.productHandler.DeleteProduct)
	api.engine.POST("/products", api.productHandler.CreateProduct)
	api.engine.PATCH("/products", api.productHandler.UpdateProduct)

	api.engine.GET("/stores/:id", api.storeHandler.GetStore)
	api.engine.DELETE("/stores/:id", api.storeHandler.DeleteStore)
	api.engine.POST("/stores", api.storeHandler.CreateStore)
	api.engine.PATCH("/stores", api.storeHandler.UpdateStore)
}
