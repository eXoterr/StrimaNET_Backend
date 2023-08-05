package app

import (
	"github.com/eXoterr/StrimaNET_Backend/internal/config"
	"github.com/eXoterr/StrimaNET_Backend/internal/database"
	"github.com/eXoterr/StrimaNET_Backend/internal/router"
)

type App struct {
	config   *config.AppConfig
	router   router.Router
	database database.Database
}

func (app *App) Run() error {
	e := app.initServices()
	if e != nil {
		return e
	}

	e = app.runServices()
	if e != nil {
		return e
	}

	return nil
}

func New() (*App, error) {
	// Try to read config.toml
	cfg := config.New()
	err := cfg.Init()
	if err != nil {
		return nil, err
	}

	return &App{config: cfg}, nil
}
