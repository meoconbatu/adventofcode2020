package day17

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point4D struct {
	x, y, z, w int
}

// Part2 func
func Part2() {
	inputFileName := "day17/input.txt"
	ins := readInput2(inputFileName)
	output := part2Core(ins)
	fmt.Println(output)
}

func part2Core(ins map[point4D]int) int {
	return countActiveCube2(ins, 6)
}
func countActiveCube2(cubes map[point4D]int, cycleth int) int {
	cube2s := make(map[point4D]int)
	for cycleth > 0 {
		cube2s = make(map[point4D]int)
		for cube := range cubes {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					for k := -1; k <= 1; k++ {
						for l := -1; l <= 1; l++ {
							c := point4D{cube.x + i, cube.y + j, cube.z + k, cube.w + l}
							num := countActiveCubeAround2(c, cubes)
							_, ok := cubes[c]
							if (ok && (num == 2 || num == 3)) || (!ok && num == 3) {
								cube2s[c] = 1
							}
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
func countActiveCubeAround2(cube point4D, cubes map[point4D]int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					c := point4D{cube.x + i, cube.y + j, cube.z + k, cube.w + l}
					if _, ok := cubes[c]; ok {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func readInput2(inputFileName string) map[point4D]int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make(map[point4D]int)
	i := 0
	for scanner.Scan() {
		for j, c := range scanner.Text() {
			if c == '#' {
				p := point4D{i, j, 0, 0}
				ins[p] = 1
			}
		}
		i++
	}
	return ins
}
