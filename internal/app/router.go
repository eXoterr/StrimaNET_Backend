package app

import "github.com/eXoterr/StrimaNET_Backend/internal/router"

func (app *App) runRouter() error {

	err := app.router.Listen()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) initRouter() error {
	app.router = router.New()
	err := app.router.Init(
		app.config.HTTP.ListenAddress,
		app.config.HTTP.Debug,
	)
	if err != nil {
		return err
	}
	return nil
}
