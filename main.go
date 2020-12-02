package main

import "flag"

func main() {
	dayth := flag.Int("day", 1, "day of puzzle")
	part := flag.Int("part", 1, "1 or 2")

	flag.Parse()

	switch *dayth {
	case 1:
		if *part == 1 {
			day1Part1()
		} else {
			day1Part2()
		}
	}
}
