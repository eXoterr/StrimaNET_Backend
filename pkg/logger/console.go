package logger

import (
	"bufio"
	"fmt"
	"time"
)

type ConsoleLogger struct {
}

func (cl *ConsoleLogger) LogData(data LoggingData, output LoggerOutput) error {

	formattedData := cl.fmtLog(data.Badge, data.Data)

	writer := bufio.NewWriter(output)

	_, err := writer.Write([]byte(formattedData))

	if err != nil {
		return err
	}

	err = writer.Flush()

	if err != nil {
		return err
	}

	return nil
}

func (cl *ConsoleLogger) fmtLog(msgType string, message string) string {
	date := time.Now()
	dateStr := date.Format("2006-1-2 15:4:5")

	msg := fmt.Sprintf("[%s] [%s] %s\n", dateStr, msgType, message)
	return msg
}
