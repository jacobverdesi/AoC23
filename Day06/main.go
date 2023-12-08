package main

import (
	"AoC23/utils"
	"fmt"
	"strconv"
	"strings"
)

func part1() {
	lines := utils.GetLinesFromFile("Day06/input.txt")
	times := strings.Fields(lines[0][6:])

	distances := strings.Fields(lines[1][9:])
	fmt.Println(times, distances)
	numWaysRecord := 1
	for idx, timeStr := range times {
		time, _ := strconv.Atoi(timeStr)
		recordDistance, _ := strconv.Atoi(distances[idx])
		numWays := 0
		for timePressed := 0; timePressed <= time; timePressed++ {
			distance := timePressed * (time - timePressed)

			if distance > recordDistance {
				numWays++
			}
		}
		numWaysRecord *= numWays
	}
	fmt.Println(numWaysRecord)

}
func part2() {
	lines := utils.GetLinesFromFile("Day06/input.txt")
	time, _ := strconv.Atoi(strings.Join(strings.Fields(lines[0][6:]), ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(lines[1][9:]), ""))
	minTime := (distance / time) + 417250 // good guess
	maxTime := time - minTime
	numWays := maxTime - minTime + 1
	fmt.Println(minTime * (time - minTime))
	fmt.Println(maxTime * (time - maxTime))
	fmt.Println(distance)
	fmt.Println(numWays)

}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
