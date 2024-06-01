package gallery_api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/data/gallery_db"
	"go-mercury/internal/service/gallery_service"
	"go-mercury/pkg/constant"
)

type GalleryHandler struct {
	galleryService *gallery_service.Service
}

func NewGalleryHandler(galleryService *gallery_service.Service) *GalleryHandler {
	return &GalleryHandler{
		galleryService: galleryService,
	}
}

func (h *GalleryHandler) getProduct(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		createFailResponse(c, http.StatusBadRequest, ErrorInvalidParam)
		return
	}

	product, err := h.galleryService.GetProduct(id)
	if err != nil {
		if errors.Is(err, constant.ProductNotFoundError) {
			createFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.GetProduct")
		createFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	createSuccessResponse(c, product)
}

func (h *GalleryHandler) createProduct(c *gin.Context) {
	reqBody, err := parseRequestBody[CreateProductRequest](c)
	if err != nil {
		createFailResponse(c, http.StatusBadRequest, err)
		return
	}

	newProduct := gallery_db.Product{
		Title:       reqBody.Title,
		ImageUrl:    reqBody.ImageUrl,
		Description: reqBody.Description,
	}

	id, err := h.galleryService.CreateProduct(newProduct)
	if err != nil {
		log.WithError(err).Error("galleryService.CreateProduct")
		createFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	createSuccessResponse(c, CreateProductResponse{Id: id})
}

func (h *GalleryHandler) updateProduct(c *gin.Context) {
	reqBody, err := parseRequestBody[UpdateProductRequest](c)
	if err != nil {
		createFailResponse(c, http.StatusBadRequest, err)
		return
	}

	product := gallery_db.Product{
		Id:          int64(reqBody.Id),
		Title:       reqBody.Title,
		ImageUrl:    reqBody.ImageUrl,
		Description: reqBody.Description,
	}

	_, err = h.galleryService.UpdateProduct(product)
	if err != nil {
		if errors.Is(err, constant.ProductNotFoundError) {
			createFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.UpdateProduct")
		createFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	createSuccessResponse(c, nil)
}
