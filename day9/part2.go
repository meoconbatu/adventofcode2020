package day9

import (
	"fmt"
	"math"
)

// Part2 func
func Part2() {
	inputFileName := "day9/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins)
	fmt.Println(output)
}
func part2Core(ins []int) int {
	target := findFirstInvalidNumber(ins)
	i, j := findSubSet(ins, target)
	min, max := findMinMax(ins[i : j+1])
	return min + max
}

// findSubSet func find a contiguous set of at least two numbers in ins array which sum to target.
//
// Return first and last index of subset found. Return 0, 0 if cannot find any subset.
func findSubSet(ins []int, target int) (int, int) {
	i, j := 0, 1
	sum := ins[i]
	for i < len(ins) && j < len(ins) {
		if sum > target {
			sum -= ins[i]
			i++
		}
		if sum == target {
			return i, j
		}
		if sum < target {
			sum += ins[j]
			j++
		}
	}
	return 0, 0
}

// findMinMax func return min and max number of ins array
func findMinMax(ins []int) (int, int) {
	min, max := math.MaxInt64, math.MinInt64
	for i := 0; i < len(ins); i++ {
		if ins[i] > max {
			max = ins[i]
		}
		if ins[i] < min {
			min = ins[i]
		}
	}
	return min, max
}
