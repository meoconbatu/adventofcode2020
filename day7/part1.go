package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type bag struct {
	num   int
	color string
}
type rule struct {
	parent   bag
	children map[string]bag
}

// Part1 func
func Part1() {
	inputFileName := "day7/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}
func part1Core(ins map[string]rule) int {
	return countBagColour(ins)
}
func countBagColour(rules map[string]rule) int {
	bags := make([]string, 1)
	bags[0] = "shiny gold"
	visited := make(map[string]bool)

	visited[bags[0]] = true
	i := 0
	for i < len(bags) {
		for _, r := range rules {
			if _, ok := r.children[bags[i]]; ok {
				if visited[r.parent.color] {
					continue
				}
				bags = append(bags, r.parent.color)
				visited[r.parent.color] = true
			}
		}
		i++
	}
	return len(visited) - 1
}

func readInput(inputFileName string) map[string]rule {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	rules := make(map[string]rule)
	for scanner.Scan() {
		re := regexp.MustCompile("([a-z ]+) bags contain([a-z0-9,\\. ]+)")

		subs := re.FindStringSubmatch(scanner.Text())

		r := rule{}
		b := bag{1, subs[1]}
		r.parent = b
		r.children = make(map[string]bag)

		re1 := regexp.MustCompile(" (\\d+) ([a-z ]+) (bags|bag)")
		bagstr := re1.FindAllStringSubmatch(subs[2], -1)
		for i := 0; i < len(bagstr); i++ {
			num, err := strconv.Atoi(bagstr[i][1])
			if err != nil {
				log.Fatalf("Want an integer, got:%s", bagstr[i][1])
			}
			tempb := bag{num, bagstr[i][2]}
			r.children[tempb.color] = tempb
		}
		rules[r.parent.color] = r
	}
	return rules
}
