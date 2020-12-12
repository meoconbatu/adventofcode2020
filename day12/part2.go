package day12

import (
	"fmt"
	"math"
)

type position2 struct {
	x, y         int
	facex, facey int
}

// Part2 func
func Part2() {
	inputFileName := "day12/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins)
	fmt.Println(output)
}
func part2Core(ins []instruction) int {
	pos := moves2(ins)
	return int(math.Abs(float64(pos.x))) + int(math.Abs(float64(pos.y)))
}
func moves2(ins []instruction) position2 {
	currentPos := position2{0, 0, 10, 1}
	for _, inst := range ins {
		move2(&currentPos, inst)
	}
	return currentPos
}
func move2(currentPos *position2, inst instruction) {
	switch inst.action {
	case forward:
		currentPos.x += currentPos.facex * inst.value
		currentPos.y += currentPos.facey * inst.value
	case north:
		currentPos.facey += inst.value
	case south:
		currentPos.facey -= inst.value
	case east:
		currentPos.facex += inst.value
	case west:
		currentPos.facex -= inst.value
	case left:
		for i := 90; i <= inst.value; i += 90 {
			currentPos.facex, currentPos.facey = -currentPos.facey, currentPos.facex
		}
	case right:
		for i := 90; i <= inst.value; i += 90 {
			currentPos.facex, currentPos.facey = currentPos.facey, -currentPos.facex
		}
	}

}
