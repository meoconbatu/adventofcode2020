package day3

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day3/input.txt"
	policies := readInput(inputFileName)
	output := part2Core(policies)
	fmt.Println(output)
}
func part2Core(ins [][]int) int {
	slopes := []struct{ y, x int }{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	result := 1
	for _, slope := range slopes {
		result *= countTree(ins, slope.y, slope.x)
	}
	return result
}
