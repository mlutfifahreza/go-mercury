package main

import (
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/api/gallery_api"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	api := gallery_api.API{}
	err := api.Run()

	if err != nil {
		log.WithError(err).Error("api.Run")
	}
}
