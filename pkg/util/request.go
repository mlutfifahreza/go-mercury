package util

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ParseRequestBody[T any](c *gin.Context) (*T, error) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(body)

	var requestBody T
	err = json.Unmarshal([]byte(bodyString), &requestBody)
	if err != nil {
		return nil, err
	}

	validate := validator.New()

	err = validate.Struct(requestBody)
	if err != nil {
		return nil, err
	}

	return &requestBody, nil
}
