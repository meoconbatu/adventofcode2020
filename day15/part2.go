package day15

import (
	"fmt"
)

type memory2 struct {
	address string
	value   int
}

// Part2 func
func Part2() {
	inputFileName := "day15/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins)
	fmt.Println(output)
}

func part2Core(ins []int) int {
	return playGame(ins, 30000000)
}
