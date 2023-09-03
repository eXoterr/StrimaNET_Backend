package logger

import (
	"fmt"

	"github.com/eXoterr/StrimaNET_Backend/pkg/logger/output/console"
	"github.com/eXoterr/StrimaNET_Backend/pkg/logger/output/file"
)

type Logger interface {
	Error(data LoggingData)
	Info(data LoggingData)
	Warning(data LoggingData)
	Debug(data LoggingData)
	SetLogLevel(level uint) error
}

type LoggingData struct {
	Badge string
	Data  string
}

var instance Logger

func GetLogger() Logger {
	if instance != nil {
		return instance
	}

	output, err := initLoggerOutput()

	if err != nil {
		fmt.Println(err)
		fmt.Printf("Unable to init logger, using default \n")
		output, _ = console.New() // Default output implementation
	}

	instance = &AbstractLogger{
		level:  0,
		impl:   &ConsoleLogger{}, // Current logger implementation
		output: output,
	}

	return instance
}

func initLoggerOutput() (LoggerOutput, error) {
	out, err := file.New() // Current output implementation
	if err != nil {
		return nil, err
	}
	return out, nil
}
