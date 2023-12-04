package main

import (
	"AoC23/utils"
	"fmt"
	"github.com/fatih/color"
	"index/suffixarray"
	"strconv"
	"strings"
)

func part1() {
	lines := utils.GetLinesFromFile("Day03/input.txt")
	lineWidth := len(lines[0])
	adjacencyOffsets := []int{-lineWidth - 1, -lineWidth, -lineWidth + 1, -1, 1, lineWidth - 1, lineWidth, lineWidth + 1}
	partNumSum := 0
	symbols := []string{"*", "%", "#", "@", "+", "-", "=", "/", "//", "&", "$"}
	symbolIdxs := make([]int, 0)
	lineIndex := suffixarray.New([]byte(strings.Join(lines, "")))
	for _, symbol := range symbols {
		symbolIdxs = append(symbolIdxs, lineIndex.Lookup([]byte(symbol), -1)...)
	}

	for idx, line := range lines {

		parts := strings.FieldsFunc(line, func(r rune) bool {
			ifSymbol := false
			for _, s := range symbols {
				if r == []rune(s)[0] {
					ifSymbol = true
					break
				}
			}
			return ifSymbol || r == '.'
		})
		prevPartIdx := 0
		for _, part := range parts {
			partIdx := strings.Index(line[prevPartIdx:], part) + (idx * lineWidth) + prevPartIdx
			partIdxEnd := partIdx + len(part)
			partHasSymbol := false
			for _, offset := range adjacencyOffsets {
				for _, symbolIdx := range symbolIdxs {
					if symbolIdx >= partIdx+offset && symbolIdx <= partIdxEnd+offset-1 {
						partHasSymbol = true
						//fmt.Println(part, partIdx, partIdxEnd, symbolIdx, offset)
						// replace partIdx->partIdxEnd with color

						break

					}
				}
				if partHasSymbol {
					break
				}
			}
			// Print between parts
			if prevPartIdx != partIdx-(idx*lineWidth) && partIdx-(idx*lineWidth) < lineWidth {
				fmt.Print(line[prevPartIdx : partIdx-(idx*lineWidth)])

			}
			prevPartIdx = partIdxEnd - (idx * lineWidth)

			// Print part
			if partHasSymbol {
				val, err := strconv.Atoi(part)
				if err != nil {
					continue
				}
				c := color.New(color.FgCyan)
				_, err = c.Printf("%s", part)
				if err != nil {
					return
				}
				partNumSum += val
			} else {
				c2 := color.New(color.FgRed)
				_, err := c2.Printf("%s", part)
				if err != nil {
					return
				}

			}
		}
		for i := prevPartIdx; i < lineWidth; i++ {
			fmt.Print(string(line[i]))
		}
		fmt.Println()

	}

	fmt.Println(partNumSum)
}
func part2() {
	lines := utils.GetLinesFromFile("Day03/input.txt")
	lineWidth := len(lines[0])
	//adjacencyOffsets := []int{-lineWidth - 1, -lineWidth, -lineWidth + 1, -1, 1, lineWidth - 1, lineWidth, lineWidth + 1}
	partNumSum := 0
	symbols := []string{"*", "%", "#", "@", "+", "-", "=", "/", "//", "&", "$"}
	partIdxs := make([]int, 0)
	lineIndex := suffixarray.New([]byte(strings.Join(lines, "")))
	gearIdxs := lineIndex.Lookup([]byte("*"), -1)

	for idx, line := range lines {

		parts := strings.FieldsFunc(line, func(r rune) bool {
			ifSymbol := false
			for _, s := range symbols {
				if r == []rune(s)[0] {
					ifSymbol = true
					break
				}
			}
			return ifSymbol || r == '.'
		})
		prevPartIdx := 0
		for _, part := range parts {
			partIdx := strings.Index(line[prevPartIdx:], part) + (idx * lineWidth) + prevPartIdx
			partIdxEnd := partIdx + len(part)
			partIdxs = append(partIdxs, partIdx)
			prevPartIdx = partIdxEnd - (idx * lineWidth)
		}
	}
	for _, gearIdx := range gearIdxs {
		adjacentParts := make([]int, 0)
		for _, partIdx := range partIdxs {
			for _, offset := range []int{-lineWidth - 1, -lineWidth, -lineWidth + 1, -1, 1, lineWidth - 1, lineWidth, lineWidth + 1} {

				if gearIdx >= partIdx+offset && gearIdx <= partIdx+offset-1 {
					fmt.Println(gearIdx, partIdx, offset)
					adjacentParts = append(adjacentParts, partIdx)
				}
			}
		}
		if len(adjacentParts) == 2 {
			partNumSum += adjacentParts[0] * adjacentParts[1]
		}
	}
	//fmt.Println(partIdxs)
	//	partIdxEnd := partIdx + len(part)
	//	partHasSymbol := false
	//	for _, offset := range adjacencyOffsets {
	//		for _, symbolIdx := range symbolIdxs {
	//			if symbolIdx >= partIdx+offset && symbolIdx <= partIdxEnd+offset-1 {
	//				partHasSymbol = true
	//				//fmt.Println(part, partIdx, partIdxEnd, symbolIdx, offset)
	//				// replace partIdx->partIdxEnd with color
	//
	//				break
	//
	//			}
	//		}
	//		if partHasSymbol {
	//			break
	//		}
	//	}
	//	// Print between parts
	//	if prevPartIdx != partIdx-(idx*lineWidth) && partIdx-(idx*lineWidth) < lineWidth {
	//		fmt.Print(line[prevPartIdx : partIdx-(idx*lineWidth)])
	//
	//	}
	//	prevPartIdx = partIdxEnd - (idx * lineWidth)
	//
	//	// Print part
	//	if partHasSymbol {
	//		val, err := strconv.Atoi(part)
	//		if err != nil {
	//			continue
	//		}
	//		c := color.New(color.FgCyan)
	//		_, err = c.Printf("%s", part)
	//		if err != nil {
	//			return
	//		}
	//		partNumSum += val
	//	} else {
	//		c2 := color.New(color.FgRed)
	//		_, err := c2.Printf("%s", part)
	//		if err != nil {
	//			return
	//		}
	//
	//	}
	//}
	//for i := prevPartIdx; i < lineWidth; i++ {
	//	fmt.Print(string(line[i]))
	//}
	//fmt.Println()
	//}

	fmt.Println(partNumSum)
}

func main() {

	//fmt.Println("--- Part 1 ---")
	//part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
