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

// Get all prime factors of a given number n
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

type Set map[int]struct{}

func (s Set) Add(item int) {
	s[item] = struct{}{}
}

func (s Set) Remove(item int) {
	delete(s, item)
}

func (s Set) Contains(item int) bool {
	_, exists := s[item]
	return exists
}
func (s Set) Iterator() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for element := range s {
			ch <- element
		}
	}()

	return ch
}
