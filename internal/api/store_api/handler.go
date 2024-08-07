package store_api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/api/general"
	"go-mercury/internal/data/gallery_db"
	"go-mercury/internal/service/gallery_service"
	"go-mercury/pkg/constant"
	"go-mercury/pkg/util"
)

type StoreHandler struct {
	galleryService *gallery_service.Service
}

func NewStoreHandler(galleryService *gallery_service.Service) *StoreHandler {
	return &StoreHandler{
		galleryService: galleryService,
	}
}

func (h *StoreHandler) GetStore(c *gin.Context) {
	id := c.Param("id")
	store, err := h.galleryService.GetStore(id)
	if err != nil {
		if errors.Is(err, constant.StoreNotFoundError) {
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.GetStore")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, store)
}

func (h *StoreHandler) DeleteStore(c *gin.Context) {
	id := c.Param("id")

	_, err := h.galleryService.DeleteStore(id)
	if err != nil {
		if errors.Is(err, constant.StoreNotFoundError) {
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.DeleteStore")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, nil)
}

func (h *StoreHandler) CreateStore(c *gin.Context) {
	reqBody, err := util.ParseRequestBody[CreateStoreRequest](c)
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, err)
		return
	}

	newStore := gallery_db.Store{
		Id:    reqBody.Id,
		Name:  reqBody.Name,
		Icon:  reqBody.Icon,
		Color: reqBody.Color,
	}

	err = h.galleryService.CreateStore(newStore)
	if err != nil {
		log.WithError(err).Error("galleryService.CreateStore")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, nil)
}

func (h *StoreHandler) UpdateStore(c *gin.Context) {
	reqBody, err := util.ParseRequestBody[UpdateStoreRequest](c)
	if err != nil {
		general.CreateFailResponse(c, http.StatusBadRequest, err)
		return
	}

	store := gallery_db.Store{
		Id:    reqBody.Id,
		Name:  reqBody.Name,
		Icon:  reqBody.Icon,
		Color: reqBody.Color,
	}

	_, err = h.galleryService.UpdateStore(store)
	if err != nil {
		if errors.Is(err, constant.StoreNotFoundError) {
			general.CreateFailResponse(c, http.StatusNotFound, err)
			return
		}

		log.WithError(err).Error("galleryService.UpdateStore")
		general.CreateFailResponse(c, http.StatusInternalServerError, err)
		return
	}

	general.CreateSuccessResponse(c, nil)
}
