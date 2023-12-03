package main

import (
	"AoC23/utils"
	"fmt"
	"strings"
	"unicode"
)

func part1() {
	lines := utils.GetLinesFromFile("Day01/input.txt")

	sum := 0
	for _, line := range lines {
		digits := make([]int, 0)
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, int(char)-48)
			}
		}
		if len(digits) != 0 {
			sum += digits[0]*10 + digits[len(digits)-1]
		}
	}
	fmt.Println(sum)
}
func part2() {
	lines := utils.GetLinesFromFile("Day01/input.txt")
	sum := 0
	for _, line := range lines {
		replacedLine := replaceNumberWords(line)

		digits := make([]int, 0)
		for _, char := range replacedLine {
			if unicode.IsDigit(char) {
				digits = append(digits, int(char)-48)
			}
		}
		lineValue := 0
		if len(digits) != 0 {
			lineValue = digits[0]*10 + digits[len(digits)-1]
			sum += lineValue
		}
		//fmt.Printf("%s %s %d %d \n", line, replacedLine, lineValue, sum)

	}
	fmt.Println(sum)

}

func replaceNumberWords(line string) string {
	minIdx := 10000
	replaceValue := ""
	replaceWordLength := 0
	for idx, word := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		idxWord := strings.Index(line, word)
		if idxWord != -1 && (idxWord < minIdx) {
			minIdx = idxWord
			replaceValue = string(rune(idx + 1 + 48))
			replaceWordLength = len(word) - 1
		}
	}
	if minIdx != 10000 {

		line = line[:minIdx] + replaceValue + line[minIdx+replaceWordLength:]
		return replaceNumberWords(line)
	}
	return line
}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
