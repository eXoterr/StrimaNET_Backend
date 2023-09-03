package logger

import (
	"encoding/json"
	"time"
)

type JSONLogger struct {
}

type jsonLog struct {
	Message string `json:"message"`
	Time    string `json:"time"`
	MsgType string `json:"type"`
}

func (jl *JSONLogger) LogData(data LoggingData, output LoggerOutput) error {

	formattedData, e := jl.fmtLog(data)

	if e != nil {
		return e
	}

	_, err := output.Write([]byte(formattedData))

	if err != nil {
		return err
	}

	return nil
}

func (jl *JSONLogger) fmtLog(data LoggingData) (string, error) {
	date := time.Now()
	dateStr := date.Format("2006-1-2 15:4:5")

	msgStruct := &jsonLog{
		Message: data.Data,
		Time:    dateStr,
		MsgType: data.Badge,
	}

	msg, e := json.Marshal(msgStruct)
	if e != nil {
		return "", e
	}
	return string(msg), nil
}
