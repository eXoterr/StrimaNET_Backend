package app

import "github.com/eXoterr/StrimaNET_Backend/pkg/logger"

func (app *App) initLogger() error {
	// Init Logging system
	logger := logger.GetLogger()
	err := logger.SetLogLevel(app.config.Logger.LoggingLevel)
	if err != nil {
		return err
	}

	return nil
}

// Implement if change logging system
// (currently no need to "start" it)
func (app *App) runLogger() error {
	return nil
}
