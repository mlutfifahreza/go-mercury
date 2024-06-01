package gallery_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func createSuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
	})
}

func createFailResponse(c *gin.Context, httpStatus int, err error) {
	c.JSON(httpStatus, APIResponse{
		Success: false,
		Error:   err.Error(),
	})
}

type CreateProductRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=256"`
	ImageUrl    string `json:"image_url" validate:"required,url"`
	Description string `json:"description" validate:"required,min=8,max=512"`
}

type CreateProductResponse struct {
	Id int `json:"id"`
}

type UpdateProductRequest struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,min=3,max=256"`
	ImageUrl    string `json:"image_url" validate:"required,url"`
	Description string `json:"description" validate:"required,min=8,max=512"`
}
