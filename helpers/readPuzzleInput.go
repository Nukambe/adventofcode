package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadPuzzleInput(path string) ([]string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	var data []byte
	data, err = os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	text := string(data)

	return strings.Split(text, "\n"), nil
}
