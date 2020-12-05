package day5

import (
	"fmt"
	"sort"
)

// Part2 func
func Part2() {
	inputFileName := "day5/input.txt"
	boardingPasses := readInput(inputFileName)
	output := part2Core(boardingPasses)
	fmt.Println(output)
}
func part2Core(ins []string) int {
	return getMissingSeatID(ins)
}
func getMissingSeatID(boardingPasses []string) int {
	seatIDs := getListSeatID(boardingPasses)
	sort.Ints(seatIDs)

	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i]+1 == seatIDs[i+1]-1 {
			return seatIDs[i] + 1
		}
	}
	return 0
}
func getListSeatID(boardingPasses []string) []int {
	seatIDs := make([]int, len(boardingPasses))
	for i, bp := range boardingPasses {
		seatIDs[i] = getSeatID(bp)
	}
	return seatIDs
}
