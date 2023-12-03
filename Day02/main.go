package main

import (
	"AoC23/utils"
	"fmt"
	"strconv"
	"strings"
)

func part1() {
	lines := utils.GetLinesFromFile("Day02/input.txt")
	sum := 0

	for idx, line := range lines {
		// split line into fields
		stripedLine := strings.ReplaceAll(line, "Game", "")
		stripedLine = strings.ReplaceAll(stripedLine, ":", "")
		stripedLine = strings.ReplaceAll(stripedLine, ",", "")
		stripedLine = strings.ReplaceAll(stripedLine, ";", " ")
		words := strings.FieldsFunc(stripedLine, func(r rune) bool {
			return r == ' '
		})[1:]
		sum += idx + 1
		for i := 0; i < len(words); i += 2 {
			//fmt.Printf("%s %s\n", words[i], words[i+1])
			val, err := strconv.Atoi(words[i])
			if err != nil {
				continue
			}

			if (words[i+1] == "red" && val > 12) || (words[i+1] == "green" && val > 13) || (words[i+1] == "blue" && val > 14) {
				sum -= idx + 1
				break
			}
		}

	}
	fmt.Println(sum)
}
func part2() {
	lines := utils.GetLinesFromFile("Day02/input.txt")
	sum := 0

	for _, line := range lines {
		// split line into fields
		stripedLine := strings.ReplaceAll(line, "Game", "")
		stripedLine = strings.ReplaceAll(stripedLine, ":", "")
		stripedLine = strings.ReplaceAll(stripedLine, ",", "")
		stripedLine = strings.ReplaceAll(stripedLine, ";", " ")
		words := strings.FieldsFunc(stripedLine, func(r rune) bool {
			return r == ' '
		})[1:]
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for i := 0; i < len(words); i += 2 {
			val, err := strconv.Atoi(words[i])
			if err != nil {
				continue
			}

			switch words[i+1] {
			case "red":
				if val > maxRed {
					maxRed = val
				}
			case "green":
				if val > maxGreen {
					maxGreen = val
				}
			case "blue":
				if val > maxBlue {
					maxBlue = val
				}
			}
		}
		sum += maxRed * maxGreen * maxBlue

	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
