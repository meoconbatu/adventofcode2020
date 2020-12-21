package day19

import (
	"fmt"
	"regexp"
	"strings"
)

// Part2 func
func Part2() {
	inputFileName := "day19/input.txt"
	ins, insr, s := readInput(inputFileName)
	output := part2Core(ins, insr, s)
	fmt.Println(output)
}

func part2Core(ins map[int][][]int, insr map[int]rune, s string) int {
	sum := 0
	regexStr42 := buildRegex(ins, insr, 42)
	regexStr31 := buildRegex(ins, insr, 31)

	regexStr := fmt.Sprintf("(?m)^(?P<first>%s{1,})(?P<second>%s{1,})$", regexStr42, regexStr31)
	re := regexp.MustCompile(regexStr)

	for _, ss := range strings.Split(s, "\n") {
		if re.FindString(ss) != ss {
			continue
		}
		match := re.FindStringSubmatch(ss)

		re42 := regexp.MustCompile("(?m)" + regexStr42)
		re31 := regexp.MustCompile("(?m)" + regexStr31)
		if len(re42.FindAllString(match[1], -1)) > len(re31.FindAllString(match[2], -1)) {
			sum++
		}
	}
	return sum
}
