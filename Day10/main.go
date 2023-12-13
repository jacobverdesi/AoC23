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
				endDirection = "E"
			case "F":
				startDirection = "S"
				endDirection = "W"
			default:
				startDirection = ""
				endDirection = ""
			}
			pipeMap[i*len(lines)+j] = Pipe{
				Index:          0,
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
	for idx, pipe := range pipeMap {
		fmt.Printf("%3d", pipe.Distance)
		if (idx+1)%len(lines) == 0 {
			fmt.Println()
		}
	}
	pipeQueue := make([]Pipe, 0)
	pipeQueue = append(pipeQueue, pipeMap[startIndex])
	for len(pipeQueue) > 0 {
		currentPipe := pipeQueue[0]
		pipeQueue = pipeQueue[1:]
		nextPipeIndex := -1
		// Assume current Pipe direction is correct
		if currentPipe.Value == "S" {

		}
		switch currentPipe.EndDirection {
		case "N":
			nextPipeIndex = currentPipe.Index - len(lines)
		case "E":
			nextPipeIndex = currentPipe.Index + 1
		case "S":
			nextPipeIndex = currentPipe.Index + len(lines)
		case "W":
			nextPipeIndex = currentPipe.Index - 1

		}
		// CHECK if next pipe matches current pipe

		nextPipe := pipeMap[nextPipeIndex]
		if nextPipe.StartDirection == currentPipe.EndDirection {
			// if connected, add to queue
			currentPipe.Distance = nextPipe.Distance + 1

			pipeQueue = append(pipeQueue, currentPipe)
		} else if nextPipe.EndDirection == currentPipe.EndDirection {
			// if connected, add to queue
			currentPipe.Distance = nextPipe.Distance + 1
			// Swap directions
			nextPipe.EndDirection, nextPipe.StartDirection = currentPipe.StartDirection, currentPipe.EndDirection
			pipeQueue = append(pipeQueue, currentPipe)
		}

	}
	for idx, pipe := range pipeMap {
		fmt.Printf("%3d", pipe.Distance)
		if (idx+1)%len(lines) == 0 {
			fmt.Println()
		}
	}
	fmt.Println(pipeMap)
}
func part2() {

}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
