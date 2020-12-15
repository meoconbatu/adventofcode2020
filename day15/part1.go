package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Part1 func
func Part1() {
	inputFileName := "day15/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}

func part1Core(ins []int) int {
	return playGame(ins, 2020)
}
func playGame(ins []int, numth int) int {
	memory := make(map[int][]int)
	for i := 0; i < len(ins); i++ {
		memory[ins[i]] = append(memory[ins[i]], i)
	}
	for len(ins) < numth {
		current := ins[len(ins)-1]
		next := 0
		if pos := memory[current]; len(pos) > 1 {
			next = pos[len(pos)-1] - pos[len(pos)-2]
		}
		ins = append(ins, next)
		memory[next] = append(memory[next], len(ins)-1)
	}
	return ins[len(ins)-1]
}

func readInput(inputFileName string) []int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	scanner.Scan()

	ins := make([]int, 0)

	for _, c := range strings.Split(scanner.Text(), ",") {
		in := 0
		fmt.Sscanf(c, "%d", &in)
		ins = append(ins, in)
	}
	return ins
}
