package gallery_api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/api/general"
	"go-mercury/internal/data/gallery_db"
	"go-mercury/internal/service/gallery_service"
	"go-mercury/pkg/constant"
	"go-mercury/pkg/util"
)

type GalleryHandler struct {
	galleryService *gallery_service.Service
}

func NewGalleryHandler(galleryService *gallery_service.Service) *GalleryHandler {
	return &GalleryHandler{
		galleryService: galleryService,
	}
}

func (h *GalleryHandler) GetProduct(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, constant.ErrorInvalidParam)
		return
	}

	product, err := h.galleryService.GetProduct(id)
	if err != nil {
		if errors.Is(err, constant.ProductNotFoundError) {
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.GetProduct")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, product)
}

func (h *GalleryHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, constant.ErrorInvalidParam)
		return
	}

	_, err = h.galleryService.DeleteProduct(id)
	if err != nil {
		if errors.Is(err, constant.ProductNotFoundError) {
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.DeleteProduct")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, nil)
}

func (h *GalleryHandler) CreateProduct(c *gin.Context) {
	reqBody, err := util.ParseRequestBody[CreateProductRequest](c)
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, err)
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
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, CreateProductResponse{Id: id})
}

func (h *GalleryHandler) UpdateProduct(c *gin.Context) {
	reqBody, err := util.ParseRequestBody[UpdateProductRequest](c)
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, err)
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
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.UpdateProduct")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, nil)
}
