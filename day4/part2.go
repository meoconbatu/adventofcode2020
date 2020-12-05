package day4

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// Part2 func
func Part2() {
	inputFileName := "day4/input.txt"
	passwords := readInput(inputFileName)
	output := part2Core(passwords)
	fmt.Println(output)
}
func part2Core(ins []password) int {
	return countValidPassword2(ins)
}
func countValidPassword2(passwords []password) int {
	numValidPass := 0
	for _, p := range passwords {
		if isValidPassword2(p) {
			numValidPass++
		}
	}
	return numValidPass
}

func isValidPassword2(pwd password) bool {
	return isValidPassword(pwd) && checkValidInt(pwd.Byr, 1920, 2002) &&
		checkValidInt(pwd.Iyr, 2010, 2020) && checkValidInt(pwd.Eyr, 2020, 2030) &&
		checkHeight(pwd.Hgt) && checkStringFormat(pwd.Hcl, "#[a-f0-9]{6}") &&
		checkStringFormat(pwd.Ecl, "^(amb|blu|brn|gry|grn|hzl|oth){1}$") &&
		checkStringFormat(pwd.Pid, "^[0-9]{9}$")
}
func checkStringFormat(field string, format string) bool {
	valid, err := regexp.MatchString(format, field)
	if err != nil {
		log.Fatal(err)
	}
	return valid
}
func checkHeight(hgt string) bool {
	num, unit := 0, ""
	fmt.Sscanf(hgt, "%d%s", &num, &unit)
	switch unit {
	case "cm":
		if !checkValidInt(strconv.Itoa(num), 150, 193) {
			return false
		}
	case "in":
		if !checkValidInt(strconv.Itoa(num), 59, 76) {
			return false
		}
	default:
		return false
	}
	return true
}
func checkValidInt(s string, min, max int) bool {
	valInt, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if valInt < min || valInt > max {
		return false
	}
	return true
}
