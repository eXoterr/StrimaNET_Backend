package app

func (app *App) initServices() error {
	err := app.initLogger() // Logger
	if err != nil {
		return err
	}

	err = app.initDatabase() // Database
	if err != nil {
		return err
	}

	err = app.initRouter() // Router
	if err != nil {
		return err
	}

	return nil
}

func (app *App) runServices() error {
	err := app.runLogger() // Logger
	if err != nil {
		return err
	}

	err = app.runDatabase() // Database
	if err != nil {
		return err
	}

	err = app.runRouter() // Router
	if err != nil {
		return err
	}

	return nil
}
