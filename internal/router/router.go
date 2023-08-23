package router

import (
	"errors"
	"io"

	"github.com/eXoterr/StrimaNET_Backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Init(addr string, debug bool) error
	Listen() error
}

type HTTPRouter struct {
	ListenAddress string
	instance      *gin.Engine
}

func (http *HTTPRouter) Listen() error {
	log := logger.GetLogger()
	if len(http.ListenAddress) == 0 {
		return errors.New("listen address in empty")
	}

	log.Debug(logger.LoggingData{
		Data: "running http server at " + http.ListenAddress,
	})
	err := http.instance.Run(http.ListenAddress)
	return err
}

func (http *HTTPRouter) Init(addr string, debug bool) error {
	log := logger.GetLogger()

	if len(addr) == 0 {
		return errors.New("listen address cannot be empty")
	}

	http.ListenAddress = addr // Set listen address

	gin.DefaultWriter = io.Discard // Write default logs to /dev/null

	if debug { // Set "mode" for router (debug/release)
		gin.SetMode(gin.DebugMode)
		log.Warning(logger.LoggingData{
			Data: "http server will run in debug mode",
		})
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.Debug(logger.LoggingData{
			Data: "http server will run in release mode",
		})
	}

	http.instance = gin.New() // Create router instance

	return nil
}

func New() Router {
	return &HTTPRouter{}
}
