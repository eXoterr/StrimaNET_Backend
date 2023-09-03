package logger

import (
	"errors"
	"io"
)

type AbstractLogger struct {
	level  uint
	impl   LoggerImpl
	output LoggerOutput
}

type LoggerImpl interface {
	LogData(data LoggingData, out LoggerOutput) error
}

type LoggerOutput io.Writer

func (ll *AbstractLogger) Error(data LoggingData) {
	if ll.level < 1 {
		return
	}

	data.Badge = "ERROR"

	ll.impl.LogData(data, ll.output)
}

func (ll *AbstractLogger) Info(data LoggingData) {
	if ll.level < 3 {
		return
	}

	data.Badge = "INFO"

	ll.impl.LogData(data, ll.output)
}

func (ll *AbstractLogger) Warning(data LoggingData) {
	if ll.level < 2 {
		return
	}

	data.Badge = "WARNING"

	ll.impl.LogData(data, ll.output)
}

func (ll *AbstractLogger) Debug(data LoggingData) {
	if ll.level < 4 {
		return
	}

	data.Badge = "DEBUG"

	ll.impl.LogData(data, ll.output)
}

func (ll *AbstractLogger) SetLogLevel(level uint) error {
	if ll.level != 0 {
		return errors.New("log level was already set")
	}

	if level == 0 {
		return errors.New("log level must be greater than 0")
	}

	if level > 4 {
		return errors.New("log level must be lower or equal 4")
	}

	ll.level = level
	return nil
}
