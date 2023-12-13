package main

import (
	"AoC23/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func findNextSequence(sequence []int) []int {
	newSequence := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		newSequence[i] = sequence[i+1] - sequence[i]
	}
	return newSequence

}
func sequenceIsAllZeros(sequence []int) bool {
	for _, val := range sequence {
		if val != 0 {
			return false
		}
	}
	return true
}
func part1() {
	lines := utils.GetLinesFromFile("Day09/input.txt")
	sumExtrapolatedVals := 0
	for _, line := range lines {
		vals := strings.Split(line, " ")
		sequence := make([]int, 0)
		for _, val := range vals {
			num, _ := strconv.Atoi(val)
			sequence = append(sequence, num)
		}
		//fmt.Println(sequence)
		nextSequence := sequence
		lastSequenceVals := make([]int, 0)
		lastSequenceVals = append(lastSequenceVals, sequence[len(sequence)-1])
		for !sequenceIsAllZeros(nextSequence) || len(nextSequence) == 0 {
			nextSequence = findNextSequence(nextSequence)
			lastSequenceVals = append(lastSequenceVals, nextSequence[len(nextSequence)-1])
			//fmt.Println(lastSequenceVals)
		}
		slices.Reverse(lastSequenceVals)
		extrapolatedVal := lastSequenceVals[0]
		for i := 0; i < len(lastSequenceVals)-1; i++ {
			extrapolatedVal += lastSequenceVals[i+1]
		}
		//fmt.Println(extrapolatedVal)
		sumExtrapolatedVals += extrapolatedVal
	}
	fmt.Println(sumExtrapolatedVals)
}
func part2() {
	lines := utils.GetLinesFromFile("Day09/input.txt")
	sumExtrapolatedVals := 0
	for _, line := range lines {
		vals := strings.Split(line, " ")
		sequence := make([]int, 0)
		for _, val := range vals {
			num, _ := strconv.Atoi(val)
			sequence = append(sequence, num)
		}
		nextSequence := sequence
		firstSequenceVals := make([]int, 0)
		firstSequenceVals = append(firstSequenceVals, sequence[0])
		for !sequenceIsAllZeros(nextSequence) || len(nextSequence) == 0 {
			nextSequence = findNextSequence(nextSequence)
			firstSequenceVals = append(firstSequenceVals, nextSequence[0])
		}
		slices.Reverse(firstSequenceVals)
		//fmt.Println(firstSequenceVals)
		extrapolatedVal := firstSequenceVals[0]
		for i := 0; i < len(firstSequenceVals)-1; i++ {
			extrapolatedVal = firstSequenceVals[i+1] - extrapolatedVal
		}
		//fmt.Println(extrapolatedVal)
		sumExtrapolatedVals += extrapolatedVal
	}
	fmt.Println(sumExtrapolatedVals)
}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
