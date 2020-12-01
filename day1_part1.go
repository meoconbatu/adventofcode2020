package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day1Part1() {
	inputFileName := "day1_input.txt"
	ins := day1ReadInput(inputFileName)
	output := day1Part1Core(ins, 2020)
	fmt.Println(output)
}

func day1Part1Core(ins []int, target int) int {
	num1, num2 := findTwo(ins, 2020)
	return num1 * num2
}

// findTwo returns two numbers in ins array that sum to target.
// Returns two 0 if cannot find the answer.
func findTwo(ins []int, target int) (int, int) {
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
func day1ReadInput(inputFileName string) []int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	ins := make([]int, 0)
	for scanner.Scan() {
		in, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Invalid input: %s", scanner.Text())
		}
		ins = append(ins, in)
	}
	return ins
}
