package gallery_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool  `json:"success,omitempty"`
	Data    any   `json:"data,omitempty"`
	Error   error `json:"error,omitempty"`
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
		Error:   err,
	})
}
