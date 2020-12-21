package day18

import (
	"fmt"
	"strconv"
)

// Part2 func
func Part2() {
	inputFileName := "day18/input.txt"
	expressions := readInput(inputFileName)
	output := part2Core(expressions)
	fmt.Println(output)
}

func part2Core(expressions []string) int {
	sum := 0
	for _, e := range expressions {
		sum += calc2(e)
	}
	return sum
}

func calc2(expression string) int {
	nums := make([]int, 0)
	opts := make([]rune, 0)
	for _, e := range expression {
		if num, err := strconv.Atoi(string(e)); err == nil {
			nums = append(nums, num)
		} else if e != ' ' {
			if len(opts) == 0 {
				opts = append(opts, e)
				continue
			}
			switch e {
			case '*':
				for len((opts)) > 0 && (opts)[len((opts))-1] == '+' {
					popOpt(&opts, &nums)
				}
				opts = append(opts, e)
			case '+':
				opts = append(opts, e)
			case '(':
				opts = append(opts, e)
			case ')':
				for opts[len(opts)-1] != '(' {
					popOpt(&opts, &nums)
				}
				opts = opts[0 : len(opts)-1]
			}
		}
	}
	for len(opts) > 0 {
		popOpt(&opts, &nums)
	}
	return nums[0]
}
