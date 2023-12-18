package main

import (
	"AoC23/utils"
	"fmt"
)

type Pipe struct {
	Index          int
	Distance       int
	StartDirection string
	EndDirection   string
	Value          string
}

func oppositeDirection(direction string) string {
	switch direction {
	case "N":
		return "S"
	case "E":
		return "W"
	case "S":
		return "N"
	case "W":
		return "E"
	}
	return ""
}
func part1() {
	lines := utils.GetLinesFromFile("Day10/input.txt")
	pipeMap := make([]Pipe, len(lines[0])*len(lines))
	for i, line := range lines {
		for j, char := range line {
			startDirection := ""
			endDirection := ""
			switch string(char) {
			case "|":
				startDirection = "N"
				endDirection = "S"
			case "-":
				startDirection = "W"
				endDirection = "E"
			case "L":
				startDirection = "N"
				endDirection = "E"
			case "J":
				startDirection = "N"
				endDirection = "W"
			case "7":
				startDirection = "S"
				endDirection = "W"
			case "F":
				startDirection = "S"
				endDirection = "E"
			default:
				startDirection = ""
				endDirection = ""
			}
			pipeMap[i*len(lines)+j] = Pipe{
				Index:          i*len(lines) + j,
				Distance:       -1,
				StartDirection: startDirection,
				EndDirection:   endDirection,
				Value:          string(char),
			}
		}
	}
	startIndex := -1
	for i, pipe := range pipeMap {
		if pipe.Value == "S" {
			startIndex = i
			break
		}
	}
	pipeMap[startIndex].Distance = 0

	currentPipeIndex := startIndex
	prevEndDirection := ""
	if pipeMap[startIndex+1].StartDirection == "W" || pipeMap[startIndex+1].EndDirection == "W" {
		currentPipeIndex = startIndex + 1
		prevEndDirection = "E"
	} else if pipeMap[startIndex-1].StartDirection == "E" || pipeMap[startIndex-1].EndDirection == "E" {
		currentPipeIndex = startIndex - 1
		prevEndDirection = "W"
	} else if pipeMap[startIndex+len(lines)].StartDirection == "N" || pipeMap[startIndex+len(lines)].EndDirection == "N" {
		currentPipeIndex = startIndex + len(lines)
		prevEndDirection = "S"
	} else if pipeMap[startIndex-len(lines)].StartDirection == "S" || pipeMap[startIndex-len(lines)].EndDirection == "S" {
		currentPipeIndex = startIndex - len(lines)
		prevEndDirection = "N"
	} else {
		panic("No start pipe found")
	}
	pipeMap[currentPipeIndex].Distance = 1
	currentPipe := pipeMap[currentPipeIndex]

	for currentPipe.Value != "S" {
		nextPipeIndex := -1
		endDirection := ""
		if currentPipe.StartDirection == oppositeDirection(prevEndDirection) {
			endDirection = currentPipe.EndDirection
		} else if currentPipe.EndDirection == oppositeDirection(prevEndDirection) {
			endDirection = currentPipe.StartDirection
		} else {
			panic("No direction found")
		}
		// CHECK if next pipe matches current pipe
		switch endDirection {
		case "N":
			nextPipeIndex = currentPipe.Index - len(lines)
		case "E":
			nextPipeIndex = currentPipe.Index + 1
		case "S":
			nextPipeIndex = currentPipe.Index + len(lines)
		case "W":
			nextPipeIndex = currentPipe.Index - 1
		}
		prevEndDirection = endDirection
		nextPipe := pipeMap[nextPipeIndex]
		pipeMap[nextPipeIndex].Distance = pipeMap[currentPipe.Index].Distance + 1
		currentPipe = nextPipe
	}

	for idx, pipe := range pipeMap {
		fmt.Printf("%10d", pipe.Distance)
		if (idx+1)%len(lines) == 0 {
			fmt.Println()
		}
	}
	fmt.Println(pipeMap[currentPipe.Index].Distance / 2)
}
func part2() {

}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
