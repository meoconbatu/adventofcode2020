package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type password struct {
	Byr string `validate:"required"`
	Iyr string `validate:"required"`
	Eyr string `validate:"required"`
	Hgt string `validate:"required"`
	Hcl string `validate:"required"`
	Ecl string `validate:"required"`
	Pid string `validate:"required"`
	Cid string
}

// Part1 func
func Part1() {
	inputFileName := "day4/input.txt"
	passwords := readInput(inputFileName)
	output := part1Core(passwords)
	fmt.Println(output)
}
func part1Core(ins []password) int {
	return countValidPassword(ins)
}
func countValidPassword(passwords []password) int {
	numValidPass := 0
	for _, p := range passwords {
		if isValidPassword(p) {
			numValidPass++
		}
	}
	return numValidPass
}
func isValidPassword(pwd password) bool {
	validate := validator.New()
	err := validate.Struct(pwd)
	return err == nil
}

func readInput(inputFileName string) []password {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	passwords := make([]password, 0)
	pass := password{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			passwords = append(passwords, pass)
			pass = password{}
			continue
		}
		for _, pair := range strings.Split(scanner.Text(), " ") {
			key, val := "", ""
			fmt.Sscanf(pair, "%3s:%s", &key, &val)

			rv := reflect.ValueOf(&pass).Elem()

			fv := rv.FieldByName(strings.Title(key))
			if !fv.IsValid() {
				log.Fatalf("Invalid input: %s", pair)
			}
			fv.SetString(val)
		}
	}
	passwords = append(passwords, pass)
	return passwords
}
