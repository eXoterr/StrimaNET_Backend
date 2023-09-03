package file

import (
	"io"
	"os"
)

func New() (io.Writer, error) {
	file, err := os.OpenFile("journal.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
