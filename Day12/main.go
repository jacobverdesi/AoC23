package main

import (
	"AoC23/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func recursiveReplace(inputString string, startIdx int, replaceVals []int) {
	replaceVal := replaceVals[0]
	for i := startIdx; i < len(inputString)-replaceVal; i++ {
		if inputString[i] == '?' {
			replacedString := inputString[:i] + "#" + inputString[i+replaceVal:]
			fmt.Println(replacedString)
			recursiveReplace(replacedString, i+1, replaceVals[1:], regex, sumArrangements)

		}

	}

}
func part1() {
	lines := utils.GetLinesFromFile("Day12/input.txt")
	sumArrangements := 0
	for _, line := range lines {
		pattern, groupsJoined := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		regex := strings.ReplaceAll(pattern, ".", "\\.")
		regex = strings.ReplaceAll(regex, "?", ".")
		re := regexp.MustCompile(regex)

		groups := strings.Split(groupsJoined, ",")
		groupInts := make([]int, len(groups))
		for i, group := range groups {
			groupInts[i], _ = strconv.Atoi(group)
		}
		recursiveReplace(pattern, 0, groupInts, re, sumArrangements)
		break
	}
	fmt.Println(sumArrangements)

}
func part2() {

}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
