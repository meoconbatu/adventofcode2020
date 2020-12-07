package main

import (
	"flag"
	"fmt"
	"meoconbatu/adventofcode2020/day1"
	"meoconbatu/adventofcode2020/day2"
	"meoconbatu/adventofcode2020/day3"
	"meoconbatu/adventofcode2020/day4"
	"meoconbatu/adventofcode2020/day5"
	"meoconbatu/adventofcode2020/day6"
)

var funcMap map[string]interface{}

func init() {
	funcMap = map[string]interface{}{
		"011": day1.Part1, "012": day1.Part2,
		"021": day2.Part1, "022": day2.Part2,
		"031": day3.Part1, "032": day3.Part2,
		"041": day4.Part1, "042": day4.Part2,
		"051": day5.Part1, "052": day5.Part2,
		"061": day6.Part1, "062": day6.Part2,
	}
}
func main() {
	dayth := flag.Int("day", 1, "day of puzzle")
	part := flag.Int("part", 1, "1 or 2")

	flag.Parse()

	execute(*dayth, *part)
}
func execute(dayth, part int) {
	key := fmt.Sprintf("%02d%d", dayth, part)
	funcMap[key].(func())()
}
