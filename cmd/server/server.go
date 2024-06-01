package main

import (
	log "github.com/sirupsen/logrus"

	"go-mercury/internal/api"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	app := api.NewApp()
	err := app.Run()

	if err != nil {
		log.WithError(err).Error("api.Run")
	}
}
