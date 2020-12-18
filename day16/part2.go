package day16

import (
	"fmt"
	"meoconbatu/adventofcode2020/utils"
)

// Part2 func
func Part2() {
	inputFileName := "day16/input.txt"
	rules, tickets := readInput(inputFileName)
	output := part2Core(tickets, rules)
	fmt.Println(output)
}

func part2Core(tickets [][]int, rules [][2]rangeType) int {
	return scanErrorRate2(tickets, rules)
}

func scanErrorRate2(tickets [][]int, rules [][2]rangeType) int {
	fieldRulesMap := make(map[int][]int)
	for i := 0; i < len(tickets); i++ {
		if isValidTicket(tickets[i], rules) {
			getRulesEligile(tickets[i], rules, fieldRulesMap)
		}
	}

	var ruleFieldMap map[int]int
	for i := 0; i < len(fieldRulesMap[0]); i++ {
		ruleFieldMap = make(map[int]int)
		if mappingFieldRule(fieldRulesMap, i, len(fieldRulesMap)-1, ruleFieldMap) {
			break
		}
	}

	result := 1
	for rule, field := range ruleFieldMap {
		if rule < 6 {
			result *= tickets[0][field]
		}
	}
	return result
}
func isValidTicket(ticket []int, rules [][2]rangeType) bool {
	for _, field := range ticket {
		if !isValidField(field, rules) {
			return false
		}
	}
	return true
}
func mappingFieldRule(m map[int][]int, x, j int, r map[int]int) bool {
	if _, ok := r[m[j][x]]; ok {
		return false
	}
	r[m[j][x]] = j
	if len(r) == len(m) {
		return true
	}
	for i := 0; i < len(m[j-1]); i++ {
		if j-1 >= 0 {
			if mappingFieldRule(m, i, j-1, r) {
				return true
			}
		}
	}
	delete(r, m[j][x])
	return false
}
func getRulesEligile(ticket []int, rules [][2]rangeType, result map[int][]int) {
	for ifield, field := range ticket {
		irule := make([]int, 0)
		for i, rule := range rules {
			if (field >= rule[0].min && field <= rule[0].max) || (field >= rule[1].min && field <= rule[1].max) {
				irule = append(irule, i)
			}
		}
		if v, ok := result[ifield]; !ok {
			result[ifield] = irule
		} else {
			result[ifield] = utils.FindIntersection(irule, v)
		}
	}
}
