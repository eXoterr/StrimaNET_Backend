package logger

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

	instance = &AbstractLogger{
		level: 0,
		impl:  &ConsoleLogger{}, // Current logger implementation
	}

	return instance
}
