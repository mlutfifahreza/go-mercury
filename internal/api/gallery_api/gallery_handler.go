package gallery_api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/service/gallery_service"
	"go-mercury/pkg/constant"
)

type GalleryHandler struct {
	galleryService *gallery_service.Service
}

func NewGalleryHandler(galleryService *gallery_service.Service) *GalleryHandler {
	return &GalleryHandler{galleryService: galleryService}
}

func (h *GalleryHandler) getPoll(c *gin.Context) {

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

		log.Errorf("h.galleryService.GetProduct(id), err=%v", err)
		createFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	createSuccessResponse(c, product)
}
