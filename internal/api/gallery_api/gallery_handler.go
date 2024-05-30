package gallery_api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-mercury/internal/service/gallery_service"
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
	}

	product, err := h.galleryService.GetProduct(id)
	if err != nil {
		createFailResponse(c, http.StatusInternalServerError, err)
	}

	createSuccessResponse(c, product)
}
