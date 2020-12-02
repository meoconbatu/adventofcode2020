package day2

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day2/input.txt"
	policies := readInput(inputFileName)
	output := part2Core(policies)
	fmt.Println(output)
}
func part2Core(policies []passwordPolicy) int {
	return countValidPassword2(policies)
}

func countValidPassword2(policies []passwordPolicy) int {
	numValidPass := 0

	for _, p := range policies {
		if isValidPassword2(p) {
			numValidPass++
		}
	}
	return numValidPass
}
func isValidPassword2(p passwordPolicy) bool {
	if (rune(p.password[p.min-1]) == p.char && rune(p.password[p.max-1]) != p.char) ||
		(rune(p.password[p.min-1]) != p.char && rune(p.password[p.max-1]) == p.char) {
		return true
	}
	return false
}
