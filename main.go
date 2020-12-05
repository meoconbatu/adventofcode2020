package main

import (
	"flag"
	"meoconbatu/adventofcode2020/day1"
	"meoconbatu/adventofcode2020/day2"
	"meoconbatu/adventofcode2020/day3"
	"meoconbatu/adventofcode2020/day4"
	"meoconbatu/adventofcode2020/day5"
)

func main() {
	dayth := flag.Int("day", 1, "day of puzzle")
	part := flag.Int("part", 1, "1 or 2")

	flag.Parse()

	switch *dayth {
	case 1:
		if *part == 1 {
			day1.Part1()
		} else {
			day1.Part2()
		}
	case 2:
		if *part == 1 {
			day2.Part1()
		} else {
			day2.Part2()
		}
	case 3:
		if *part == 1 {
			day3.Part1()
		} else {
			day3.Part2()
		}
	case 4:
		if *part == 1 {
			day4.Part1()
		} else {
			day4.Part2()
		}
	case 5:
		if *part == 1 {
			day5.Part1()
		} else {
			day5.Part2()
		}
	}

}
