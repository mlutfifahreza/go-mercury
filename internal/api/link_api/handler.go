package link_api

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

type LinkHandler struct {
	galleryService *gallery_service.Service
}

func NewLinkHandler(galleryService *gallery_service.Service) *LinkHandler {
	return &LinkHandler{
		galleryService: galleryService,
	}
}

func (h *LinkHandler) GetLink(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, constant.ErrorInvalidParam)
		return
	}

	Link, err := h.galleryService.GetLink(id)
	if err != nil {
		if errors.Is(err, constant.LinkNotFoundError) {
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.GetLink")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, Link)
}

func (h *LinkHandler) DeleteLink(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, constant.ErrorInvalidParam)
		return
	}

	_, err = h.galleryService.DeleteLink(id)
	if err != nil {
		if errors.Is(err, constant.LinkNotFoundError) {
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.DeleteLink")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, nil)
}

func (h *LinkHandler) CreateLink(c *gin.Context) {
	reqBody, err := util.ParseRequestBody[CreateLinkRequest](c)
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, err)
		return
	}

	newLink := gallery_db.Link{
		ProductId: reqBody.ProductId,
		StoreId:   reqBody.StoreId,
		Link:      reqBody.Link,
	}

	id, err := h.galleryService.CreateLink(newLink)
	if err != nil {
		if errors.Is(err, constant.StoreNotFoundError) || errors.Is(err, constant.StoreNotFoundError) {
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.CreateLink")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, CreateLinkResponse{Id: id})
}

func (h *LinkHandler) UpdateLink(c *gin.Context) {
	reqBody, err := util.ParseRequestBody[UpdateLinkRequest](c)
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, err)
		return
	}

	Link := gallery_db.Link{
		Id:        reqBody.Id,
		ProductId: reqBody.ProductId,
		StoreId:   reqBody.StoreId,
		Link:      reqBody.Link,
	}

	_, err = h.galleryService.UpdateLink(Link)
	if err != nil {
		if errors.Is(err, constant.LinkNotFoundError) {
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.UpdateLink")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, nil)
}
