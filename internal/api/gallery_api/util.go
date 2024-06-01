package gallery_api

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func parseRequestBody[T any](c *gin.Context) (*T, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(body)

	var requestBpdy T
	err = json.Unmarshal([]byte(bodyString), &requestBpdy)
	if err != nil {
		return nil, err
	}

	validate := validator.New()

	err = validate.Struct(requestBpdy)
	if err != nil {
		return nil, err
	}

	return &requestBpdy, nil
}
