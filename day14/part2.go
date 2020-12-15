package day14

import (
	"fmt"
)

type memory2 struct {
	address string
	value   int
}

// Part2 func
func Part2() {
	inputFileName := "day14/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins)
	fmt.Println(output)
}

func part2Core(ins []input) int {
	return sumMemory2(ins)
}
func sumMemory2(ins []input) int {
	mems := []memory2{}
	for _, in := range ins {
		mems = append(mems, applyBitmask2(in)...)
	}
	addrs := make(map[string]int)
	for i := len(mems) - 1; i >= 0; i-- {
		generateAddr(mems[i], "", &addrs)
	}
	sum := 0
	for _, val := range addrs {
		sum += val
	}
	return sum
}
func applyBitmask2(in input) []memory2 {
	result := []memory2{}
	for _, mem := range in.mems {
		addr := fmt.Sprintf("%036b", mem.address)
		for i, m := range in.bitmask {
			switch m {
			case 'X':
				addr = addr[0:i] + "X" + addr[i+1:]
			case '1':
				addr = addr[0:i] + "1" + addr[i+1:]
			}
		}
		result = append(result, memory2{addr, mem.value})
	}
	return result
}
func generateAddr(mem memory2, addr string, result *map[string]int) {
	if len(mem.address) == 0 {
		if _, ok := (*result)[addr]; !ok {
			(*result)[addr] = mem.value
		}
		return
	}
	s := mem.address[0]
	mem.address = mem.address[1:]
	if s == 'X' {
		generateAddr(mem, addr+"0", result)
		generateAddr(mem, addr+"1", result)
	} else {
		generateAddr(mem, addr+string(s), result)
	}
}
