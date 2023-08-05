package logger

import (
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"
)

type ConsoleLogger struct {
	logLevel uint
}

func (cl *ConsoleLogger) Info(data string) {
	if cl.logLevel < 3 {
		return
	}
	badge := color.BlueString("INFO")
	msg := cl.fmtLog(badge, data)

	fmt.Println(msg)
}

func (cl *ConsoleLogger) Error(data string) {
	if cl.logLevel < 1 {
		return
	}
	badge := color.RedString("ERROR")
	msg := cl.fmtLog(badge, data)

	fmt.Println(msg)
}

func (cl *ConsoleLogger) Warning(data string) {
	if cl.logLevel < 2 {
		return
	}
	badge := color.YellowString("WARN")
	msg := cl.fmtLog(badge, data)

	fmt.Println(msg)
}

func (cl *ConsoleLogger) Debug(data string) {
	if cl.logLevel < 4 {
		return
	}

	badge := color.WhiteString("DEBUG")
	msg := cl.fmtLog(badge, data)

	fmt.Println(msg)
}

func (cl *ConsoleLogger) fmtLog(msgType string, message string) string {
	date := time.Now()
	dateStr := date.Format("2006-1-2 15:4:5")

	msg := fmt.Sprintf("[%s] [%s] %s", dateStr, msgType, message)
	return msg
}

func (cl *ConsoleLogger) SetLogLevel(level uint) error {
	if cl.logLevel != 0 {
		return errors.New("log level was already set")
	}

	if level == 0 {
		return errors.New("log level must be greater than 0")
	}

	if level > 4 {
		return errors.New("log level must be lower or equal 4")
	}

	cl.logLevel = level

	return nil
}
