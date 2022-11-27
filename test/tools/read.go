package tools

import (
	"io"
	"os"
)

func ReadFile(path string) (string, error) {
	fs, err := os.Open(path)
	if err != nil {
		return "", err
	}
	content, err := io.ReadAll(fs)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
