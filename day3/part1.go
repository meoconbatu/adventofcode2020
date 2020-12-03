package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Part1 func
func Part1() {
	inputFileName := "day3/input.txt"
	policies := readInput(inputFileName)
	output := part1Core(policies)
	fmt.Println(output)
}
func part1Core(ins [][]int) int {
	return countTree(ins, 3, 1)
}
func countTree(ins [][]int, y, x int) int {
	numTree := 0
	i, j := 0, 0
	for {
		j += y
		if j >= len(ins[i]) {
			j -= len(ins[i])
		}
		i += x
		if i >= len(ins) {
			break
		}
		if ins[i][j] == 1 {
			numTree++
		}
	}
	return numTree
}
func readInput(inputFileName string) [][]int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make([][]int, 0)
	for scanner.Scan() {
		in := make([]int, 0)
		for _, r := range scanner.Text() {
			val := 0
			if r == '#' {
				val = 1
			}
			in = append(in, val)
		}
		ins = append(ins, in)
	}
	return ins
}
