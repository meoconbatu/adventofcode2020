package day13

import (
	"fmt"
)

type input2 struct {
	index, busID int
}

// Part2 func
func Part2() {
	inputFileName := "day13/input.txt"
	_, busIDs := readInput(inputFileName)
	output := part2Core(busIDs)
	fmt.Println(output)
}
func part2Core(busIDs []int) int {
	ins := processInput(busIDs)
	return findEarliestTimestamp(ins)
}

// findEarliestTimestamp func uses Chinese Remainder Theorem
func findEarliestTimestamp(ins []input2) int {
	prod := 1
	rem := make([]int, len(ins))
	pp := make([]int, len(ins))
	inv := make([]int, len(ins))
	for _, in := range ins {
		prod *= in.busID
	}
	for i := 1; i < len(ins); i++ {
		rem[i] = ins[i].busID - ins[i].index
		pp[i] = prod / ins[i].busID
		inv[i] = findInverse(pp[i], ins[i].busID)
	}
	x := 0
	for i := 1; i < len(ins); i++ {
		x += rem[i] * pp[i] * inv[i]
	}
	return x % prod
}

// findInverse func return x which a*x % b = 1
func findInverse(a, b int) int {
	for x := 1; x < b; x++ {
		if a*x%b == 1 {
			return x
		}
	}
	return 0
}

// processInput func eliminates 0 elements, convert []int to []input2
func processInput(busIDs []int) []input2 {
	output := make([]input2, 0)
	for j, id := range busIDs {
		if id != 0 {
			output = append(output, input2{j, id})
		}
	}
	return output
}
