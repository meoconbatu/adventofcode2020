package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type instruction struct {
	operation string
	argument  int
	exec      bool
}

// Part1 func
func Part1() {
	inputFileName := "day8/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}
func part1Core(ins []instruction) int {
	accumulator, _ := execBootCode(ins)
	return accumulator
}
func execBootCode(ins []instruction) (int, int) {
	accumulator := 0
	i := 0
	for i < len(ins) && !ins[i].exec {
		ins[i].exec = true
		switch ins[i].operation {
		case "nop":
			i++
		case "acc":
			accumulator += ins[i].argument
			i++
		case "jmp":
			i += ins[i].argument
		}
	}
	return accumulator, i
}
func readInput(inputFileName string) []instruction {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	instructions := make([]instruction, 0)
	for scanner.Scan() {
		opt, arg := "", 0
		fmt.Sscanf(scanner.Text(), "%s %d", &opt, &arg)

		instructions = append(instructions, instruction{opt, arg, false})
	}
	return instructions
}
