package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type input struct {
	bitmask string
	mems    []memory
}
type memory struct {
	address, value int
}

// Part1 func
func Part1() {
	inputFileName := "day14/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}

func part1Core(ins []input) int {
	return sumMemory(ins)
}
func sumMemory(ins []input) int {
	result := &map[int]int{}
	for _, in := range ins {
		applyBitmask(in, result)
	}
	sum := 0
	for _, val := range *result {
		sum += val
	}
	return sum
}
func applyBitmask(in input, result *map[int]int) {
	for j := range in.mems {
		for i, m := range in.bitmask {
			switch m {
			case '0':
				in.mems[j].value &^= (1 << (len(in.bitmask) - 1 - i))
			case '1':
				in.mems[j].value |= (1 << (len(in.bitmask) - 1 - i))
			}
		}
		(*result)[in.mems[j].address] = in.mems[j].value
	}
}
func readInput(inputFileName string) []input {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make([]input, 0)
	in := input{}
	for scanner.Scan() {
		bitmask := ""
		fmt.Sscanf(scanner.Text(), "mask = %s", &bitmask)
		if bitmask != "" {
			if in.bitmask != "" {
				ins = append(ins, in)
			}
			in = input{}
			in.bitmask = bitmask
			continue
		}
		addr, val := 0, 0
		fmt.Sscanf(scanner.Text(), "mem[%d] = %d", &addr, &val)
		in.mems = append(in.mems, memory{addr, val})
	}
	ins = append(ins, in)
	return ins
}
