package day6

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day6/input.txt"
	groups := readInput(inputFileName)
	output := part2Core(groups)
	fmt.Println(output)
}
func part2Core(ins [][]string) int {
	return sumCountAnswerYes2(ins)
}
func sumCountAnswerYes2(groups [][]string) int {
	sum := 0
	for _, g := range groups {
		sum += countAnswerYes2(g)
	}
	return sum
}
func countAnswerYes2(group []string) int {
	sum := 0
	bit := 0xFFFFFFFF
	for _, questions := range group {
		bitTemp := 0
		for _, question := range questions {
			n := question - 'a'
			bitTemp |= (1 << n)
		}
		bit &= bitTemp
	}
	for bit > 0 {
		if bit&1 > 0 {
			sum++
		}
		bit >>= 1
	}
	return sum
}
