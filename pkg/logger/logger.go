package logger

type Logger interface {
	Error(data string)
	Info(data string)
	Warning(data string)
	Debug(data string)
	SetLogLevel(level uint) error
}

var instance Logger

func GetLogger() Logger {
	if instance != nil {
		return instance
	}

	instance = &ConsoleLogger{
		logLevel: 0,
	}

	return instance
}
