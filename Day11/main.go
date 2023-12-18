package main

import (
	"AoC23/utils"
	"fmt"
	"math"
)

type xy struct {
	idx int
	x   int
	y   int
}

func l1Distance(a xy, b xy) int {
	return int(math.Abs(float64(b.x-a.x))) + int(math.Abs(float64(b.y-a.y)))
}
func part1() {
	lines := utils.GetLinesFromFile("Day11/input.txt")
	emptyRowIdxs := make(utils.Set)
	emptyColIdxs := make(utils.Set)
	galaxyIdxs := make([]xy, 0)

	for i, line := range lines {
		emptyRowIdxs.Add(i)
		emptyColIdxs.Add(i)
		for j, char := range line {
			if string(char) == "#" {
				galaxyIdxs = append(galaxyIdxs, xy{idx: len(galaxyIdxs) + 1, x: j, y: i})
			}
		}
	}
	for _, galaxy := range galaxyIdxs {
		if emptyRowIdxs.Contains(galaxy.y) {
			emptyRowIdxs.Remove(galaxy.y)
		}
		if emptyColIdxs.Contains(galaxy.x) {
			emptyColIdxs.Remove(galaxy.x)
		}
	}
	//fmt.Println(emptyRowIdxs)
	//fmt.Println(emptyColIdxs)
	sumDistances := 0
	for i := 0; i < len(galaxyIdxs); i++ {
		for j := i + 1; j < len(galaxyIdxs); j++ {
			start := galaxyIdxs[i]
			end := galaxyIdxs[j]
			//fmt.Println(start, end)
			//startEnd := end
			for emptyRowIdx := range emptyRowIdxs.Iterator() {

				if galaxyIdxs[i].y < emptyRowIdx && emptyRowIdx < galaxyIdxs[j].y {
					end.y++
				} else if galaxyIdxs[i].y > emptyRowIdx && emptyRowIdx > galaxyIdxs[j].y {
					start.y++
				}
			}
			for emptyColIdx := range emptyColIdxs.Iterator() {
				if galaxyIdxs[i].x < emptyColIdx && emptyColIdx < galaxyIdxs[j].x {
					end.x++
				} else if galaxyIdxs[i].x > emptyColIdx && emptyColIdx > galaxyIdxs[j].x {
					start.x++
				}
			}
			//fmt.Println(start, end)
			//fmt.Println(start.idx, end.idx, startEnd, end, l1Distance(start, end))
			sumDistances += l1Distance(start, end)
		}
	}
	fmt.Println(sumDistances)

}
func part2() {
	lines := utils.GetLinesFromFile("Day11/input.txt")
	emptyRowIdxs := make(utils.Set)
	emptyColIdxs := make(utils.Set)
	galaxyIdxs := make([]xy, 0)

	for i, line := range lines {
		emptyRowIdxs.Add(i)
		emptyColIdxs.Add(i)
		for j, char := range line {
			if string(char) == "#" {
				galaxyIdxs = append(galaxyIdxs, xy{idx: len(galaxyIdxs) + 1, x: j, y: i})
			}
		}
	}
	for _, galaxy := range galaxyIdxs {
		if emptyRowIdxs.Contains(galaxy.y) {
			emptyRowIdxs.Remove(galaxy.y)
		}
		if emptyColIdxs.Contains(galaxy.x) {
			emptyColIdxs.Remove(galaxy.x)
		}
	}
	//fmt.Println(emptyRowIdxs)
	//fmt.Println(emptyColIdxs)
	sumDistances := 0
	distanceMultiplier := 999999
	for i := 0; i < len(galaxyIdxs); i++ {
		for j := i + 1; j < len(galaxyIdxs); j++ {
			start := galaxyIdxs[i]
			end := galaxyIdxs[j]
			//fmt.Println(start, end)
			for emptyRowIdx := range emptyRowIdxs.Iterator() {

				if galaxyIdxs[i].y < emptyRowIdx && emptyRowIdx < galaxyIdxs[j].y {
					end.y += distanceMultiplier
				} else if galaxyIdxs[i].y > emptyRowIdx && emptyRowIdx > galaxyIdxs[j].y {
					start.y += distanceMultiplier
				}
			}
			for emptyColIdx := range emptyColIdxs.Iterator() {
				if galaxyIdxs[i].x < emptyColIdx && emptyColIdx < galaxyIdxs[j].x {
					end.x += distanceMultiplier
				} else if galaxyIdxs[i].x > emptyColIdx && emptyColIdx > galaxyIdxs[j].x {
					start.x += distanceMultiplier
				}
			}
			//fmt.Println(start, end)
			sumDistances += l1Distance(start, end)
		}
	}
	fmt.Println(sumDistances)
}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
