package day9

import (
	"bufio"
	"fmt"
	"log"
	"meoconbatu/adventofcode2020/day1"
	"os"
)

// Part1 func
func Part1() {
	inputFileName := "day9/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}
func part1Core(ins []int) int {
	return findFirstInvalidNumber(ins)
}

// findFirstInvalidNumber func finds first number that is invalid.
// Valid number is a number that is the sum of any two of the 25 immediately previous numbers.
// Return 0 if cannot find any invalid number.
func findFirstInvalidNumber(ins []int) int {
	for i := 25; i < len(ins); i++ {
		x, y := day1.FindTwo(ins[i-25:i], ins[i])
		if x == 0 && y == 0 {
			return ins[i]
		}
	}
	return 0
}
func readInput(inputFileName string) []int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make([]int, 0)
	for scanner.Scan() {
		i := 0
		fmt.Sscanf(scanner.Text(), "%d", &i)

		ins = append(ins, i)
	}
	return ins
}
