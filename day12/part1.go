package day12

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	north int = iota
	east
	south
	west
	left
	right
	forward
)

type instruction struct {
	action, value int
}
type position struct {
	x, y int
	face int
}

var actionMaps map[rune]int

func init() {
	actionMaps = map[rune]int{'L': left, 'R': right, 'N': north, 'E': east, 'W': west, 'S': south, 'F': forward}
}

// Part1 func
func Part1() {
	inputFileName := "day12/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}
func part1Core(ins []instruction) int {
	pos := moves(ins)
	return int(math.Abs(float64(pos.x))) + int(math.Abs(float64(pos.y)))
}
func moves(ins []instruction) position {
	currentPos := position{0, 0, east}
	for _, inst := range ins {
		move(&currentPos, inst)
	}
	return currentPos
}
func move(currentPos *position, inst instruction) {
	newAction := inst.action
	if inst.action == forward {
		newAction = currentPos.face
	}
	switch newAction {
	case north:
		currentPos.y += inst.value
	case south:
		currentPos.y -= inst.value
	case east:
		currentPos.x += inst.value
	case west:
		currentPos.x -= inst.value
	case left:
		currentPos.face = (currentPos.face - inst.value/90 + 4) % 4
	case right:
		currentPos.face = (currentPos.face + inst.value/90) % 4
	}

}
func readInput(inputFileName string) []instruction {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make([]instruction, 0)
	for scanner.Scan() {
		action, actionInt, value := ' ', 0, 0
		fmt.Sscanf(scanner.Text(), "%c%d", &action, &value)

		if _, ok := actionMaps[action]; !ok {
			log.Fatalf("Invalid input: %c", action)
		}
		actionInt = actionMaps[action]

		in := instruction{actionInt, value}
		ins = append(ins, in)
	}
	return ins
}
