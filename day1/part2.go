package day1

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day1/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins, 2020)
	fmt.Print(output)
}

func part2Core(ins []int, target int) int {
	num1, num2, num3 := findThree(ins, 2020)
	return num1 * num2 * num3
}

// findThree returns three numbers in ins array that sum to target.
// Returns three 0 if cannot find the answer.
func findThree(ins []int, target int) (int, int, int) {
	for i := 0; i < len(ins); i++ {
		num1 := ins[i]
		num2, num3 := findTwo(ins[i:], target-num1)
		if num2 != 0 {
			return num1, num2, num3
		}
	}
	return 0, 0, 0
}
