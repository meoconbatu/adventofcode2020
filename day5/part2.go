package day5

import (
	"fmt"
	"meoconbatu/adventofcode2020/utils"
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

// getMissingSeatID func O(nlogn)
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

// getMissingSeatID2 func O(n)
func getMissingSeatID2(boardingPasses []string) int {
	seatIDs := getListSeatID(boardingPasses)
	min, max := utils.FindMinMax(seatIDs)

	xor := 0
	for i := min; i <= max; i++ {
		xor ^= i
	}

	for i := 0; i < len(seatIDs); i++ {
		xor ^= seatIDs[i]
	}
	return xor
}
func getListSeatID(boardingPasses []string) []int {
	seatIDs := make([]int, len(boardingPasses))
	for i, bp := range boardingPasses {
		seatIDs[i] = getSeatID(bp)
	}
	return seatIDs
}
