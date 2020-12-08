package day7

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day7/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins)
	fmt.Println(output)
}
func part2Core(ins map[string]rule) int {
	return countBagColour2(ins)
}

var sum int

func countBagColour2(rules map[string]rule) int {
	sum = 0
	colour := "shiny gold"
	// f(rules, bag{1, colour})
	return f(rules, bag{1, colour})
}

func f(rules map[string]rule, target bag) int {
	if len(rules[target.color].children) == 0 {
		return 0
	}
	temp := 0
	for _, k := range rules[target.color].children {
		temp += k.num + k.num*f(rules, k)
	}
	// sum += temp
	return temp
}
