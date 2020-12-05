package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Part1 func
func Part1() {
	inputFileName := "day5/input.txt"
	boardingPasses := readInput(inputFileName)
	output := part1Core(boardingPasses)
	fmt.Println(output)
}
func part1Core(ins []string) int {
	return getMaxSeatID(ins)
}
func getMaxSeatID(boardingPasses []string) int {
	maxSeatID := 0
	for _, bp := range boardingPasses {
		seatID := getSeatID(bp)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}
func getSeatID(bp string) int {
	return getRowNumber(bp[0:7], 0, 127)*8 + getColumnNumber(bp[7:], 0, 7)
}

func getRowNumber(bp string, min, max int) int {
	if bp == "" {
		return min
	}
	if rune(bp[0]) == 'F' {
		return getRowNumber(bp[1:], min, (min+max)/2)
	}
	return getRowNumber(bp[1:], (min+max)/2+1, max)
}

func getColumnNumber(bp string, min, max int) int {
	if bp == "" {
		return min
	}
	if rune(bp[0]) == 'L' {
		return getColumnNumber(bp[1:], min, (min+max)/2)
	}
	return getColumnNumber(bp[1:], (min+max)/2+1, max)
}

func readInput(inputFileName string) []string {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	boardingPasses := make([]string, 0)
	for scanner.Scan() {
		boardingPasses = append(boardingPasses, scanner.Text())
	}
	return boardingPasses
}
