package console

import (
	"fmt"
	"io"
)

type ConsoleOutput struct {
}

func (co ConsoleOutput) Write(data []byte) (int, error) {
	fmt.Println(string(data))
	return len(data), nil
}

func New() (io.Writer, error) {
	return ConsoleOutput{}, nil
}
