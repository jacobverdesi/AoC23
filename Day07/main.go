package main

import (
	"AoC23/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"unicode"
)

func getRuneFrequency(input string) []int {
	uniqueChars := make(map[rune]int)
	for _, char := range input {
		if _, exists := uniqueChars[char]; !exists {
			uniqueChars[char] = 1
		} else {
			uniqueChars[char]++
		}
	}
	freqs := make([]int, 0)
	for _, freq := range uniqueChars {
		freqs = append(freqs, freq)
	}
	slices.Sort(freqs)
	slices.Reverse(freqs)

	return freqs
}
func getRankP1(hand string) int {
	handRank := 0
	for idx, card := range hand {
		if unicode.IsDigit(card) {
			handRank += int(math.Pow(100, float64(len(hand)-idx-1))) * int(card-'0')
		} else {
			cardVal := 0
			switch card {
			case 'T':
				cardVal = 10
			case 'J':
				cardVal = 11
			case 'Q':
				cardVal = 12
			case 'K':
				cardVal = 13
			case 'A':
				cardVal = 14
			}
			handRank += int(math.Pow(100, float64(len(hand)-idx-1))) * cardVal // dont need to map to 10-14
		}
	}
	charFreqs := getRuneFrequency(hand)
	switch charFreqs[0] {
	case 5:
		handRank += 7 * utils.PowInt(10, 10)
	case 4:
		handRank += 6 * utils.PowInt(10, 10)
	case 3:
		if charFreqs[1] == 2 {
			handRank += 5 * utils.PowInt(10, 10)
		} else {
			handRank += 4 * utils.PowInt(10, 10)
		}
	case 2:
		if charFreqs[1] == 2 {
			handRank += 3 * utils.PowInt(10, 10)
		} else {
			handRank += 2 * utils.PowInt(10, 10)
		}
	default:
		handRank += 1 * utils.PowInt(10, 10)

	}
	return handRank
}
func getRankP2(hand string) int {
	handRank := 0
	for idx, card := range hand {
		if unicode.IsDigit(card) {
			handRank += int(math.Pow(100, float64(len(hand)-idx-1))) * int(card-'0')
		} else {
			cardVal := 0
			switch card {
			case 'T':
				cardVal = 10
			case 'J':
				cardVal = 0
			case 'Q':
				cardVal = 12
			case 'K':
				cardVal = 13
			case 'A':
				cardVal = 14
			}
			handRank += int(math.Pow(100, float64(len(hand)-idx-1))) * cardVal // dont need to map to 10-14
		}
	}
	charFreqs := getRuneFrequency(hand)
	for _, card := range hand {
		if card == 'J' {
			charFreqs[0]++
		}
	}
	setOffset := utils.PowInt(10, 10)
	switch charFreqs[0] {
	case 5:
		handRank += 7 * setOffset
	case 4:
		handRank += 6 * setOffset
	case 3:
		if charFreqs[1] == 2 {
			handRank += 5 * setOffset
		} else {
			handRank += 4 * setOffset
		}
	case 2:
		if charFreqs[1] == 2 {
			handRank += 3 * setOffset
		} else {
			handRank += 2 * setOffset
		}
	default:
		handRank += 1 * setOffset

	}
	return handRank
}
func part1() {
	lines := utils.GetLinesFromFile("Day07/input.txt")
	ranks := make(map[int]int)
	for _, line := range lines {
		hand := line[:5]
		bid, _ := strconv.Atoi(line[6:])
		handRank := getRankP1(hand)
		ranks[handRank] = bid
	}
	rankKeys := make([]int, 0)
	for k := range ranks {
		rankKeys = append(rankKeys, k)
	}
	slices.Sort(rankKeys)
	totalWin := 0
	for keyIdx, rankKey := range rankKeys {
		totalWin += ranks[rankKey] * (keyIdx + 1)
	}
	fmt.Println(totalWin)
}
func part2() {
	lines := utils.GetLinesFromFile("Day07/input.txt")
	ranks := make(map[int]int)
	for _, line := range lines {
		hand := line[:5]
		bid, _ := strconv.Atoi(line[6:])
		handRank := getRankP2(hand)
		ranks[handRank] = bid
	}
	rankKeys := make([]int, 0)
	for k := range ranks {
		rankKeys = append(rankKeys, k)
	}
	slices.Sort(rankKeys)
	totalWin := 0
	for keyIdx, rankKey := range rankKeys {
		totalWin += ranks[rankKey] * (keyIdx + 1)
	}
	fmt.Println(totalWin)
}

func main() {
	fmt.Println("--- Part 1 ---")
	part1()
	fmt.Println("--- Part 2 ---")
	part2()
}
