package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Part1 func
func Part1() {
	inputFileName := "day19/input.txt"
	rules, runes, s := readInput(inputFileName)
	output := part1Core(rules, runes, s)
	fmt.Println(output)
}

func part1Core(rules map[int][][]int, runes map[int]rune, s string) int {
	regexStr := buildRegex(rules, runes, 0)
	regexStr = "(?m)^" + regexStr + "$"

	re := regexp.MustCompile(regexStr)

	return len(re.FindAllString(s, -1))
}
func buildRegex(rules map[int][][]int, runes map[int]rune, rule int) (t string) {
	if _, ok := runes[rule]; ok {
		return string(runes[rule])
	}
	for _, subrules := range rules[rule] {
		temp := ""
		for _, subrule := range subrules {
			temp += buildRegex(rules, runes, subrule)
		}
		if t != "" {
			t = fmt.Sprintf("(?:%s|%s)", t, temp)
		} else {
			t = temp
		}
	}
	return
}

func readInput(inputFileName string) (map[int][][]int, map[int]rune, string) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make(map[int][][]int)
	insr := make(map[int]rune)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		k, r := 0, ' '
		fmt.Sscanf(scanner.Text(), "%d: \"%c\"", &k, &r)
		if r != ' ' {
			insr[k] = r
			continue
		}

		rule := 0
		parts := strings.Split(scanner.Text(), ":")
		fmt.Sscanf(parts[0], "%d", &rule)

		rules := make([][]int, 0)
		for _, p := range strings.Split(parts[1], "|") {
			in := make([]int, 0)
			for _, r := range strings.Split(strings.Trim(p, " "), " ") {
				val, _ := strconv.Atoi(r)
				in = append(in, val)
			}
			rules = append(rules, in)
		}
		ins[rule] = rules
	}

	s := ""
	for scanner.Scan() {
		s += scanner.Text() + "\n"
	}
	return ins, insr, strings.Trim(s, "\n")
}
