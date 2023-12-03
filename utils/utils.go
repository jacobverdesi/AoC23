package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetLinesFromFile(filename string) []string {
	absPath, _ := filepath.Abs(filename)
	body, _ := os.ReadFile(absPath)
	bodyString := string(body)

	lines := strings.FieldsFunc(bodyString, func(r rune) bool {
		return r == '\n' || r == '\r'
	})
	return lines
}
