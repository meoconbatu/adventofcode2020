package day10

import (
	"fmt"
	"sort"
)

// Part2 func
func Part2() {
	inputFileName := "day10/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins)
	fmt.Println(output)
}
func part2Core(ins []int) int {
	return countArrangements(ins)
}

func countArrangements(ins []int) int {
	m := make(map[int]int)
	m[3] = 7
	m[2] = 4
	m[1] = 2
	sort.Ints(ins)
	ins = append(ins, ins[len(ins)-1]+3)

	result := 1
	i, j := 0, 1
	for j < len(ins) {
		if ins[j-1]+1 < ins[j] {
			temp := 0
			if i == 0 {
				temp = j - i - 1
			} else {
				temp = j - i - 2
			}
			if _, ok := m[temp]; ok {
				result *= m[temp]
			}
			i = j
		}
		j++
	}
	return result
}
