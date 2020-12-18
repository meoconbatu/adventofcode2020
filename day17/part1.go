package day17

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point3D struct {
	x, y, z int
}

// Part1 func
func Part1() {
	inputFileName := "day17/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}

func part1Core(ins map[point3D]int) int {
	return countActiveCube(ins, 6)
}
func countActiveCube(cubes map[point3D]int, cycleth int) int {
	cube2s := make(map[point3D]int)
	for cycleth > 0 {
		cube2s = make(map[point3D]int)
		for cube := range cubes {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					for k := -1; k <= 1; k++ {
						c := point3D{cube.x + i, cube.y + j, cube.z + k}
						num := countActiveCubeAround(c, cubes)
						_, ok := cubes[c]
						if (ok && (num == 2 || num == 3)) || (!ok && num == 3) {
							cube2s[c] = 1
						}
					}
				}
			}
		}
		cubes, cube2s = cube2s, cubes
		cycleth--
	}
	return len(cubes)
}
func countActiveCubeAround(cube point3D, cubes map[point3D]int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				c := point3D{cube.x + i, cube.y + j, cube.z + k}
				if _, ok := cubes[c]; ok {
					sum++
				}
			}
		}
	}
	return sum
}

func readInput(inputFileName string) map[point3D]int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make(map[point3D]int)
	i := 0
	for scanner.Scan() {
		for j, c := range scanner.Text() {
			if c == '#' {
				p := point3D{i, j, 0}
				ins[p] = 1
			}
		}
		i++
	}
	return ins
}
