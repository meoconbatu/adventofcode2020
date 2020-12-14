package day13

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Part1 func
func Part1() {
	inputFileName := "day13/input.txt"
	timestamp, busIDs := readInput(inputFileName)
	output := part1Core(timestamp, busIDs)
	fmt.Println(output)
}
func part1Core(timestamp int, busIDs []int) int {
	minTime, minID := findEarliestBus(timestamp, busIDs)
	return (minTime - timestamp) * minID
}
func findEarliestBus(timestamp int, busIDs []int) (int, int) {
	minTime, minBusID := math.MaxInt64, 0
	for _, id := range busIDs {
		if id == 0 {
			continue
		}
		temp := int(float64(id) * math.Ceil(float64(timestamp)/float64(id)))
		if temp < minTime {
			minTime = temp
			minBusID = id
		}
	}
	return minTime, minBusID
}
func readInput(inputFileName string) (int, []int) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	timestamp := 0
	ins := make([]int, 0)

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &timestamp)

	scanner.Scan()
	for _, c := range strings.Split(scanner.Text(), ",") {
		in := 0
		fmt.Sscanf(c, "%d", &in)
		ins = append(ins, in)
	}

	return timestamp, ins
}
