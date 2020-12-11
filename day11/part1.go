package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	floor int = iota
	emptySeat
	occupiedSeat
)

// Part1 func
func Part1() {
	inputFileName := "day11/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}
func part1Core(ins [][]int) int {
	return countOccupiedSeats(ins, isEligibleToChangeSeat)
}

func countOccupiedSeats(ins [][]int, isEligibleToChangeSeat func([][]int, int, int) bool) int {
	result := 0
	ischanged := true

	copyins := make([][]int, len(ins))
	for i := 0; i < len(ins); i++ {
		copyins[i] = make([]int, len(ins[i]))
		for j := 0; j < len(ins[i]); j++ {
			copyins[i][j] = ins[i][j]
		}
	}

	x, y := &ins, &copyins

	for ischanged {
		ischanged = false
		for i := 0; i < len(ins); i++ {
			for j := 0; j < len(ins[i]); j++ {
				if isEligibleToChangeSeat(*x, i, j) {
					ischanged = true
					switch (*x)[i][j] {
					case emptySeat:
						(*y)[i][j] = occupiedSeat
					case occupiedSeat:
						(*y)[i][j] = emptySeat
					}
				} else {
					(*y)[i][j] = (*x)[i][j]
				}
			}
		}
		x, y = y, x
	}
	for i := 0; i < len(ins); i++ {
		for j := 0; j < len(ins[i]); j++ {
			if ins[i][j] == occupiedSeat {
				result++
			}
		}
	}

	return result
}
func isEligibleToChangeSeat(ins [][]int, i, j int) bool {
	switch (ins)[i][j] {
	case emptySeat:
		return countAdjacentOccupiedSeat(ins, i, j) == 0
	case occupiedSeat:
		return countAdjacentOccupiedSeat(ins, i, j) >= 4
	}
	return false
}
func countAdjacentOccupiedSeat(ins [][]int, i, j int) int {
	count := 0
	for ii := i - 1; ii <= i+1 && ii < len(ins); ii++ {
		for jj := j - 1; jj <= j+1 && jj < len(ins[i]); jj++ {
			if ii < 0 || jj < 0 {
				continue
			}
			if ins[ii][jj] == occupiedSeat && (ii != i || jj != j) {
				count++
			}
		}
	}
	return count
}
func readInput(inputFileName string) [][]int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make([][]int, 0)
	for scanner.Scan() {
		in := make([]int, len(scanner.Text()))
		for i, r := range scanner.Text() {
			val := 0
			switch r {
			case '.':
				val = floor
			case 'L':
				val = emptySeat
			case '#':
				val = occupiedSeat
			}
			in[i] = val
		}
		ins = append(ins, in)
	}
	return ins
}
