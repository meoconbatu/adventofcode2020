package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type passwordPolicy struct {
	min, max int
	char     rune
	password string
}

// Part1 func
func Part1() {
	inputFileName := "day2/input.txt"
	policies := readInput(inputFileName)
	output := part1Core(policies)
	fmt.Println(output)
}
func part1Core(policies []passwordPolicy) int {
	return countValidPassword(policies)
}

func countValidPassword(policies []passwordPolicy) int {
	numValidPass := 0

	for _, p := range policies {
		if isValidPassword(p) {
			numValidPass++
		}
	}
	return numValidPass
}
func isValidPassword(p passwordPolicy) bool {
	n := 0
	for _, c := range p.password {
		if c == p.char {
			n++
		}
		if n > p.max {
			return false
		}
	}
	return n >= p.min
}

// 5-9 t: ttttjtttttttttt
func readInput(inputFileName string) []passwordPolicy {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	policies := make([]passwordPolicy, 0)
	for scanner.Scan() {
		min, max := 0, 0
		char := ' '
		password := ""
		fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &min, &max, &char, &password)
		passPolicy := passwordPolicy{min, max, char, password}
		policies = append(policies, passPolicy)
	}
	return policies
}
