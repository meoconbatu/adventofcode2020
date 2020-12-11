package main

import (
	"fmt"
	"log"
	"meoconbatu/adventofcode2020/config"
	"meoconbatu/adventofcode2020/day1"
	"meoconbatu/adventofcode2020/day10"
	"meoconbatu/adventofcode2020/day11"
	"meoconbatu/adventofcode2020/day2"
	"meoconbatu/adventofcode2020/day3"
	"meoconbatu/adventofcode2020/day4"
	"meoconbatu/adventofcode2020/day5"
	"meoconbatu/adventofcode2020/day6"
	"meoconbatu/adventofcode2020/day7"
	"meoconbatu/adventofcode2020/day8"
	"meoconbatu/adventofcode2020/day9"
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
		"071": day7.Part1, "072": day7.Part2,
		"081": day8.Part1, "082": day8.Part2,
		"091": day9.Part1, "092": day9.Part2,
		"101": day10.Part1, "102": day10.Part2,
		"111": day11.Part1, "112": day11.Part2,
	}
}
func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	execute(conf.Dayth, conf.Part)
}
func execute(dayth, part int) {
	key := fmt.Sprintf("%02d%d", dayth, part)
	if f, ok := funcMap[key]; ok {
		f.(func())()
	} else {
		log.Fatalf("Invalid flag input, got:%d, %d", dayth, part)
	}
}
