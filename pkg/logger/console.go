package logger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
}

func (cl *ConsoleLogger) LogData(data LoggingData) error {

	formattedData := cl.fmtLog(data.Badge, data.Data)

	fmt.Println(formattedData)

	return nil
}

func (cl *ConsoleLogger) fmtLog(msgType string, message string) string {
	date := time.Now()
	dateStr := date.Format("2006-1-2 15:4:5")

	msg := fmt.Sprintf("[%s] [%s] %s", dateStr, msgType, message)
	return msg
}
