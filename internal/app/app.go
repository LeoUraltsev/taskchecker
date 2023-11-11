package app

import (
	"context"
	"github.com/LeoUraltsev/taskchecker/config"
	"github.com/LeoUraltsev/taskchecker/internal/models"
	"github.com/LeoUraltsev/taskchecker/internal/repo"
	"github.com/LeoUraltsev/taskchecker/pkg/postgres"
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

	pg, _ := postgres.New(context.Background(), "qwe")
	r := repo.NewRepository(pg)
	r.User.CreateUser(context.TODO(), models.User{
		UserID:   0,
		Username: "",
		Email:    "",
		Password: "",
		BIO:      "",
	})
	log.Fatal(e.StartServer(&httpserver))

}
