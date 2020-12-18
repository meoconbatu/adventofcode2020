package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type rangeType struct {
	min, max int
}

// Part1 func
func Part1() {
	inputFileName := "day16/input.txt"
	rules, tickets := readInput(inputFileName)
	output := part1Core(tickets, rules)
	fmt.Println(output)
}

func part1Core(tickets [][]int, rules [][2]rangeType) int {
	return scanErrorRate(tickets, rules)
}
func scanErrorRate(tickets [][]int, rules [][2]rangeType) int {
	sum := 0
	for i := 1; i < len(tickets); i++ {
		for _, field := range tickets[i] {
			if !isValidField(field, rules) {
				sum += field
			}
		}
	}
	return sum
}
func isValidField(field int, rules [][2]rangeType) bool {
	for _, rule := range rules {
		for _, r := range rule {
			if field >= r.min && field <= r.max {
				return true
			}
		}
	}
	return false
}

func readInput(inputFileName string) ([][2]rangeType, [][]int) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	rules := make([][2]rangeType, 0)
	tickets := make([][]int, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		field := ""
		min1, max1, min2, max2 := 0, 0, 0, 0
		fmt.Sscanf(scanner.Text(), "%s %d-%d or %d-%d", &field, &min1, &max1, &min2, &max2)
		if min1 == 0 {
			fmt.Sscanf(scanner.Text(), "%s%s %d-%d or %d-%d", &field, &field, &min1, &max1, &min2, &max2)
		}
		rules = append(rules, [2]rangeType{{min1, max1}, {min2, max2}})
	}

	scanner.Scan()

	scanner.Scan()
	ins := make([]int, 0)
	for _, c := range strings.Split(scanner.Text(), ",") {
		in := 0
		fmt.Sscanf(c, "%d", &in)
		ins = append(ins, in)
	}
	tickets = append(tickets, ins)

	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		ins := make([]int, 0)
		for _, c := range strings.Split(scanner.Text(), ",") {
			in := 0
			fmt.Sscanf(c, "%d", &in)
			ins = append(ins, in)
		}
		tickets = append(tickets, ins)
	}
	return rules, tickets
}
