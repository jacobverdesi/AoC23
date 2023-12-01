package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func part1() {
	// get relative path to input.txt
	absPath, _ := filepath.Abs("Day00/input.txt")
	body, _ := os.ReadFile(absPath)
	line := ""
	sum := 0
	firstDigit := 0
	lastDigit := 0
	for _, char := range body {
		if char == '\n' {
			continue
		}
		if char == '\r' {
			line = ""

			sum += firstDigit*10 + lastDigit
			firstDigit = 0
			lastDigit = 0
			continue
		}
		if int(char) >= 48 && int(char) <= 57 {
			if firstDigit == 0 {
				firstDigit = int(char) - 48
			}
			lastDigit = int(char) - 48

		}
		line += string(char)
	}
	fmt.Println(sum)
}
func part2() {
	// get relative path to input.txt
	absPath, _ := filepath.Abs("Day00/input.txt")
	body, _ := os.ReadFile(absPath)
	line := ""
	sum := 0
	firstDigit := 0
	lastDigit := 0
	replaced := string(body)
	replaced = strings.ReplaceAll(replaced, "one", "1")
	replaced = strings.ReplaceAll(replaced, "two", "2")
	replaced = strings.ReplaceAll(replaced, "three", "3")
	replaced = strings.ReplaceAll(replaced, "four", "4")
	replaced = strings.ReplaceAll(replaced, "five", "5")
	replaced = strings.ReplaceAll(replaced, "six", "6")
	replaced = strings.ReplaceAll(replaced, "seven", "7")
	replaced = strings.ReplaceAll(replaced, "eight", "8")
	replaced = strings.ReplaceAll(replaced, "nine", "9")

	for _, char := range replaced {
		if char == '\n' {
			continue
		}
		if char == '\r' {

			sum += firstDigit*10 + lastDigit
			firstDigit = 0
			lastDigit = 0
			line = ""

			continue
		}
		if int(char) >= 48 && int(char) <= 57 {
			if firstDigit == 0 {
				firstDigit = int(char) - 48
			}
			lastDigit = int(char) - 48

		}
		line += string(char)
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
