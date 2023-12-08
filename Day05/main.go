package main

import (
	"AoC23/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Ranges struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}
type SeedRange struct {
	seedStart int
	seedEnd   int
}

func part1() {
	lines := utils.GetLinesFromFile("Day05/input.txt")
	seeds := make([]int, 0)
	mapTypeIdx := 0
	maps := make(map[int][]Ranges)
	lowestLocationNumber := math.MaxInt64

	for lineIdx, line := range lines {
		if lineIdx == 0 {
			seedsStr := strings.Fields(line[6:])
			for _, seedStr := range seedsStr {
				seed, _ := strconv.Atoi(seedStr)
				seeds = append(seeds, seed)
			}
		} else if !unicode.IsDigit(rune(line[0])) {
			mapTypeIdx++
		} else if line == "" {
			continue
		} else {
			ranges := strings.Fields(line)
			destinationStart, _ := strconv.Atoi(ranges[0])
			sourceStart, _ := strconv.Atoi(ranges[1])
			rangeLength, _ := strconv.Atoi(ranges[2])
			maps[mapTypeIdx] = append(maps[mapTypeIdx], Ranges{destinationStart, sourceStart, rangeLength})

		}

	}
	mapKeys := make([]int, 0)
	for k := range maps {
		mapKeys = append(mapKeys, k)
	}
	sort.Ints(mapKeys)

	for _, seed := range seeds {
		prevSeedMapVal := seed
		for _, mapKey := range mapKeys {
			for _, rangeVal := range maps[mapKey] {
				if prevSeedMapVal >= rangeVal.sourceStart && prevSeedMapVal < rangeVal.sourceStart+rangeVal.rangeLength {
					prevSeedMapVal = rangeVal.destinationStart + (prevSeedMapVal - rangeVal.sourceStart)
					break
				}
			}
		}
		if prevSeedMapVal < lowestLocationNumber {
			lowestLocationNumber = prevSeedMapVal
		}
	}
	fmt.Println(lowestLocationNumber)

}
func part2() {
	lines := utils.GetLinesFromFile("Day05/input.txt")
	seeds := make([]SeedRange, 0)
	mapTypeIdx := 0
	maps := make(map[int][]Ranges)
	lowestLocationNumber := math.MaxInt64

	for lineIdx, line := range lines {
		if lineIdx == 0 {
			seedsStr := strings.Fields(line[6:])

			for i := 0; i < len(seedsStr); i += 2 {
				seedStart, _ := strconv.Atoi(seedsStr[i])
				seedRange, _ := strconv.Atoi(seedsStr[i+1])
				seeds = append(seeds, SeedRange{seedStart, seedStart + seedRange})
			}
		} else if !unicode.IsDigit(rune(line[0])) {
			mapTypeIdx++
		} else if line == "" {
			continue
		} else {
			ranges := strings.Fields(line)
			destinationStart, _ := strconv.Atoi(ranges[0])
			sourceStart, _ := strconv.Atoi(ranges[1])
			rangeLength, _ := strconv.Atoi(ranges[2])
			maps[mapTypeIdx] = append(maps[mapTypeIdx], Ranges{destinationStart, sourceStart, rangeLength})

		}

	}
	mapKeys := make([]int, 0)
	for k := range maps {
		mapKeys = append(mapKeys, k)
	}
	sort.Ints(mapKeys)

	for _, seedRange := range seeds {
		// list of seed ranges
		// for each map for each map range check if seed range is within,outside, or overlaps
		// if within, get range from destination
		// if outside, append to list of seed ranges
		// if overlaps, split seed range into two seed ranges and append to list of seed ranges
		// repeat until no more seed ranges
		// find min of seed ranges

		prevSeedRanges := make([]SeedRange, 0)
		prevSeedRanges = append(prevSeedRanges, seedRange)
		for _, mapKey := range mapKeys {
			for _, rangeVal := range maps[mapKey] {

				newSeedRanges := make([]SeedRange, 0)
				for _, prevSeedRange := range prevSeedRanges {
					if prevSeedRange.seedStart >= rangeVal.sourceStart && prevSeedRange.seedEnd <= rangeVal.sourceStart+rangeVal.rangeLength {
						// within
						newSeedRanges = append(newSeedRanges, SeedRange{rangeVal.destinationStart + (prevSeedRange.seedStart - rangeVal.sourceStart), rangeVal.destinationStart + (prevSeedRange.seedEnd - rangeVal.sourceStart)})
					} else if prevSeedRange.seedStart < rangeVal.sourceStart && prevSeedRange.seedEnd > rangeVal.sourceStart+rangeVal.rangeLength {
						// overlaps
						newSeedRanges = append(newSeedRanges, SeedRange{prevSeedRange.seedStart, rangeVal.sourceStart})
						newSeedRanges = append(newSeedRanges, SeedRange{rangeVal.destinationStart, rangeVal.rangeLength})
						newSeedRanges = append(newSeedRanges, SeedRange{rangeVal.sourceStart + rangeVal.rangeLength, prevSeedRange.seedEnd})
					} else if prevSeedRange.seedStart < rangeVal.sourceStart && prevSeedRange.seedEnd > rangeVal.sourceStart {
						// overlap start
						newSeedRanges = append(newSeedRanges, SeedRange{prevSeedRange.seedStart, rangeVal.sourceStart})
						newSeedRanges = append(newSeedRanges, SeedRange{rangeVal.destinationStart, prevSeedRange.seedEnd})
					} else if prevSeedRange.seedStart < rangeVal.sourceStart+rangeVal.rangeLength && prevSeedRange.seedEnd > rangeVal.sourceStart+rangeVal.rangeLength {
						// overlap end
						newSeedRanges = append(newSeedRanges, SeedRange{rangeVal.destinationStart + (prevSeedRange.seedStart - rangeVal.sourceStart), rangeVal.destinationStart + (rangeVal.sourceStart + rangeVal.rangeLength)})
						newSeedRanges = append(newSeedRanges, SeedRange{rangeVal.sourceStart + rangeVal.rangeLength, prevSeedRange.seedEnd})
					} else if prevSeedRange.seedStart < rangeVal.sourceStart {
						// outside
						newSeedRanges = append(newSeedRanges, prevSeedRange)
					} else if prevSeedRange.seedEnd > rangeVal.sourceStart+rangeVal.rangeLength {
						// outside
						newSeedRanges = append(newSeedRanges, prevSeedRange)
					}
				}
				prevSeedRanges = newSeedRanges
				fmt.Println(len(prevSeedRanges))
			}
		}
		for _, prevSeedRange := range prevSeedRanges {
			if prevSeedRange.seedStart < lowestLocationNumber {
				lowestLocationNumber = prevSeedRange.seedStart
			}
		}

	}
	fmt.Println(lowestLocationNumber)

}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
