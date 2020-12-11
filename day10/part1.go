package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// Part1 func
func Part1() {
	inputFileName := "day10/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}
func part1Core(ins []int) int {
	result := findJoltDiff(ins)
	return result[1] * result[3]
}

func findJoltDiff(ins []int) map[int]int {
	sort.Ints(ins)
	ins = append(ins, ins[len(ins)-1]+3)
	x := 0
	result := make(map[int]int)
	for _, in := range ins {
		y := in
		result[y-x]++
		x = in
	}
	return result
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
