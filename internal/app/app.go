package app

import (
	"github.com/LeoUraltsev/taskchecker/config"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Run() {

	// Init config
	cfg := config.New()

	// Init logger
	SetLogrus(cfg.Log.Level)

	log.Debugf("Configuration: %v\n", cfg)

	//Echo handler
	log.Info("Initializing handler and routes...")
	e := echo.New()

	//Http server
	log.Info("Starting server...")
	httpserver := http.Server{
		Addr:         cfg.Http.Port,
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(e.StartServer(&httpserver))

}
