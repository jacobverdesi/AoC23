package main

import (
	"AoC23/utils"
	"fmt"
	"slices"
)

type node struct {
	val   string
	left  string
	right string
}

func part1() {
	lines := utils.GetLinesFromFile("Day08/input.txt")
	pathInstruction := lines[0]
	nodeMap := make(map[string]node)
	for _, line := range lines[1:] {
		val := line[0:3]
		left := line[7:10]
		right := line[12:15]
		nodeMap[val] = node{val, left, right}
	}
	steps := 0
	found := false
	currentNode := nodeMap["AAA"]
	for !found {
		for _, char := range pathInstruction {
			if currentNode.val == "ZZZ" {
				found = true
				break
			}
			if string(char) == "L" {
				currentNode = nodeMap[currentNode.left]
			} else {
				currentNode = nodeMap[currentNode.right]
			}
			steps++
		}
	}
	fmt.Println("Steps:", steps)
}
func part2() {
	lines := utils.GetLinesFromFile("Day08/input.txt")
	pathInstruction := lines[0]
	nodeMap := make(map[string]node)

	for _, line := range lines[1:] {
		val := line[0:3]
		left := line[7:10]
		right := line[12:15]
		nodeMap[val] = node{val, left, right}
	}
	currentNodes := make([]node, 0)
	for node1 := range nodeMap {
		if nodeMap[node1].val[2] == 'A' {
			currentNodes = append(currentNodes, nodeMap[node1])
		}
	}
	factors := make([]int, 0)
	for _, node1 := range currentNodes {
		breakLoop := false
		nodeStep := 0
		visited := make(map[string]bool)
		for !breakLoop {
			for _, char := range pathInstruction {

				if node1.val[2] == 'Z' {
					visited[node1.val] = true
				}
				if visited[node1.val] {
					breakLoop = true
					nodeFactors := utils.PrimeFactors(nodeStep)
					for _, stepFactors := range nodeFactors {
						if !slices.Contains(factors, stepFactors) {
							factors = append(factors, stepFactors)
						}
					}
					break
				}
				if string(char) == "L" {
					node1 = nodeMap[node1.left]
				} else {
					node1 = nodeMap[node1.right]
				}
				nodeStep++
			}
		}
	}
	steps := 1
	for _, factor := range factors {
		steps *= factor
	}
	fmt.Println("Steps:", steps)
}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
