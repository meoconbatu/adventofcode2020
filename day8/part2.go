package day8

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day8/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins)
	fmt.Println(output)
}

func part2Core(ins []instruction) int {
	return execBootCode2(ins)
}
func execBootCode2(ins []instruction) int {
	iprev := -1
	for i := range ins {
		if ins[i].operation == "nop" {
			ins[i].operation = "jmp"
		} else if ins[i].operation == "jmp" {
			ins[i].operation = "nop"
		} else {
			continue
		}
		if iprev != -1 && iprev != i {
			if ins[iprev].operation == "nop" {
				ins[iprev].operation = "jmp"
			} else if ins[iprev].operation == "jmp" {
				ins[iprev].operation = "nop"
			}
		}
		iprev = i
		for ii := range ins {
			ins[ii].exec = false
		}
		acc, indx := execBootCode(ins)
		if indx == len(ins) {
			return acc
		}
	}
	return 0
}
