package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Part1 func
func Part1() {
	inputFileName := "day6/input.txt"
	groups := readInput(inputFileName)
	output := part1Core(groups)
	fmt.Println(output)
}
func part1Core(ins [][]string) int {
	return sumCountAnswerYes(ins)
}
func sumCountAnswerYes(groups [][]string) int {
	sum := 0
	for _, g := range groups {
		sum += countAnswerYes(g)
	}
	return sum
}
func countAnswerYes(group []string) int {
	sum := 0
	bit := 0
	for _, questions := range group {
		for _, question := range questions {
			n := question - 'a'
			if bit&(1<<n) == 0 {
				sum++
				bit |= (1 << n)
			}
		}
	}
	return sum
}
func readInput(inputFileName string) [][]string {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	groups := make([][]string, 0)
	group := []string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			groups = append(groups, group)
			group = []string{}
			continue
		}
		group = append(group, scanner.Text())
	}
	groups = append(groups, group)
	return groups
}
