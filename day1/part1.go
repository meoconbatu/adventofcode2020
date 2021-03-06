package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Part1 func
func Part1() {
	inputFileName := "day1/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins, 2020)
	fmt.Println(output)
}

func part1Core(ins []int, target int) int {
	num1, num2 := FindTwo(ins, 2020)
	return num1 * num2
}

// FindTwo returns two numbers in ins array that sum to target.
// Returns 0, 0 if cannot find the answer.
func FindTwo(ins []int, target int) (int, int) {
	hashTable := make(map[int]bool)
	for _, in := range ins {
		complement := target - in
		if hashTable[complement] {
			return in, complement
		}
		hashTable[in] = true
	}
	return 0, 0
}
func readInput(inputFileName string) []int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	ins := make([]int, 0)
	for scanner.Scan() {
		in := 0
		_, err := fmt.Sscanf(scanner.Text(), "%d", &in)
		if err != nil {
			log.Fatalf("Invalid input: %s", scanner.Text())
		}
		ins = append(ins, in)
	}
	return ins
}
