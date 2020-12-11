package day11

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day11/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins)
	fmt.Println(output)
}
func part2Core(ins [][]int) int {
	return countOccupiedSeats(ins, isEligibleToChangeSeat2)
}

func isEligibleToChangeSeat2(ins [][]int, i, j int) bool {
	switch (ins)[i][j] {
	case emptySeat:
		return countAdjacentOccupiedSeat2(ins, i, j) == 0
	case occupiedSeat:
		return countAdjacentOccupiedSeat2(ins, i, j) >= 5
	}
	return false
}
func countAdjacentOccupiedSeat2(ins [][]int, i, j int) int {
	count := 0
	for ii := i - 1; ii <= i+1 && ii < len(ins); ii++ {
		for jj := j - 1; jj <= j+1 && jj < len(ins[i]); jj++ {
			if ii < 0 || jj < 0 || (ii == i && jj == j) {
				continue
			}
			if (ii != i || jj != j) && findFirstSeat(ins, ii, jj, ii-i, jj-j) == occupiedSeat {
				count++
			}
		}
	}
	return count
}
func findFirstSeat(ins [][]int, i, j, stepi, stepj int) int {
	for i >= 0 && j >= 0 && i < len(ins) && j < len(ins[i]) {
		if ins[i][j] != floor {
			return ins[i][j]
		}
		i += stepi
		j += stepj
	}
	return floor
}
