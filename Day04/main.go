package main

import (
	"AoC23/utils"
	"fmt"
	"math"
	"strings"
)

func part1() {
	lines := utils.GetLinesFromFile("Day04/input.txt")
	cardScore := 0.0
	for _, line := range lines {
		splitPipe := strings.Split(line, "|")
		winningNumbers := strings.Fields(splitPipe[0])
		currentNumberSet := strings.Fields(splitPipe[1])
		numWinning := 0
		for _, currentNumber := range currentNumberSet {
			for _, winningNumber := range winningNumbers {
				if currentNumber == winningNumber {
					numWinning++
				}
			}
		}
		if numWinning != 0 {
			cardScore += math.Pow(2, float64(numWinning-1))
		}

	}
	fmt.Println(cardScore)

}

func part2() {
	lines := utils.GetLinesFromFile("Day04/input.txt")
	cardScore := 0
	cardMap := make(map[int]int)
	for numLines := range lines {
		cardMap[numLines] = 1
	}
	fmt.Println(cardMap)
	for idx, line := range lines {
		splitPipe := strings.Split(line, "|")
		winningNumbers := strings.Fields(splitPipe[0])
		currentNumberSet := strings.Fields(splitPipe[1])
		numWinning := 0
		for _, currentNumber := range currentNumberSet {
			for _, winningNumber := range winningNumbers {
				if currentNumber == winningNumber {
					numWinning++
				}
			}
		}
		for idx2 := 0; idx2 < numWinning; idx2++ {
			cardMap[idx2+idx+1] += cardMap[idx]
		}
		fmt.Println(cardMap)

	}
	for _, val := range cardMap {
		cardScore += val
	}
	fmt.Println(cardScore)
}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
