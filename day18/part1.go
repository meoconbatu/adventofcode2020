package day18

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Part1 func
func Part1() {
	inputFileName := "day18/input.txt"
	expressions := readInput(inputFileName)
	output := part1Core(expressions)
	fmt.Println(output)
}

func part1Core(expressions []string) int {
	sum := 0
	for _, e := range expressions {
		sum += calc(e)
	}
	return sum
}

func calc(expression string) int {
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
			case '+', '*':
				popOpt(&opts, &nums)
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
	if len(opts) > 0 {
		popOpt(&opts, &nums)
	}
	return nums[0]
}
func popOpt(opts *[]rune, nums *[]int) {
	if (*opts)[len((*opts))-1] == '(' {
		return
	}
	num1, num2 := (*nums)[len((*nums))-1], (*nums)[len((*nums))-2]
	(*nums) = (*nums)[0 : len((*nums))-2]
	num := 0
	switch (*opts)[len((*opts))-1] {
	case '+':
		num = num1 + num2
	case '*':
		num = num1 * num2
	}
	(*opts) = (*opts)[0 : len((*opts))-1]
	(*nums) = append((*nums), num)
}
func readInput(inputFileName string) []string {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make([]string, 0)

	for scanner.Scan() {
		ins = append(ins, scanner.Text())
	}

	return ins
}
