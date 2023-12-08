package utils

import (
	"math"
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

func FindNthOccurrence(str string, substr string, n int) int {
	idx := -1
	for i := 0; i <= n; i++ {
		idx = strings.Index(str[idx+1:], substr)
		if idx == -1 {
			return -1
		}
		idx += idx + 1
	}
	return idx
}
func IndexAll(str string, substr string) []int {
	idx := -1
	indices := make([]int, 0)
	for {
		idx = strings.Index(str[idx+1:], substr)
		if idx == -1 {
			return indices
		}
		idx += idx + 1
		indices = append(indices, idx)
	}

}
func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
func GetUniqueChars(input string) string {
	uniqueChars := make(map[rune]bool)
	result := ""

	for _, char := range input {
		if _, exists := uniqueChars[char]; !exists {
			uniqueChars[char] = true
			result += string(char)
		}
	}

	return result
}
