package app

import (
	"github.com/eXoterr/StrimaNET_Backend/internal/database"
)

func (app *App) initDatabase() error {
	app.database = database.New()
	err := app.database.Configure()

	return err
}

func (app *App) runDatabase() error {

	err := app.database.Connect()
	if err != nil {
		return err
	}

	return err
}
