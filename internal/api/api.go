package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/api/auth_api"
	"go-mercury/internal/api/link_api"
	"go-mercury/internal/api/product_api"
	"go-mercury/internal/api/store_api"
	"go-mercury/internal/data/gallery_db"
	"go-mercury/internal/service/auth_service"
	"go-mercury/internal/service/gallery_service"
)

type App struct {
	engine         *gin.Engine
	productHandler *product_api.ProductHandler
	storeHandler   *store_api.StoreHandler
	linkHandler    *link_api.LinkHandler
	authHandler    *auth_api.AuthHandler
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
	authService := auth_service.NewService(galleryDB)
	api.productHandler = product_api.NewProductHandler(&galleryService)
	api.storeHandler = store_api.NewStoreHandler(&galleryService)
	api.linkHandler = link_api.NewLinkHandler(&galleryService)
	api.authHandler = auth_api.NewAuthHandler(&authService)

	return nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (api *App) SetupRouter() {
	api.engine.GET("/ping", Ping)
	api.engine.GET("/healthcheck", HealthCheck)

	api.engine.Use(CORSMiddleware())

	api.engine.GET("/products", api.productHandler.GetProductList)
	api.engine.POST("/products", api.productHandler.CreateProduct)
	api.engine.PATCH("/products", api.productHandler.UpdateProduct)
	api.engine.GET("/products/:id", api.productHandler.GetProduct)
	api.engine.DELETE("/products/:id", api.productHandler.DeleteProduct)
	api.engine.GET("/products/:id/detail", api.productHandler.GetProductDetail)

	api.engine.GET("/stores/:id", api.storeHandler.GetStore)
	api.engine.DELETE("/stores/:id", api.storeHandler.DeleteStore)
	api.engine.POST("/stores", api.storeHandler.CreateStore)
	api.engine.PATCH("/stores", api.storeHandler.UpdateStore)

	api.engine.GET("/links/:id", api.linkHandler.GetLink)
	api.engine.DELETE("/links/:id", api.linkHandler.DeleteLink)
	api.engine.POST("/links", api.linkHandler.CreateLink)
	api.engine.PATCH("/links", api.linkHandler.UpdateLink)

	api.engine.POST("/auth/register", api.authHandler.Register)
	api.engine.POST("/auth/login", api.authHandler.Login)
	api.engine.GET("/auth/user-data", api.authHandler.UserData)
}
