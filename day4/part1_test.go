package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name string
		pwd  password
		want bool
	}{
		{"1", password{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Cid: "147", Hgt: "183cm"}, true},
		{"1", password{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Cid: "147"}, false},
		{"1", password{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Hgt: "183cm"}, true},
		{"1", password{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Iyr: "2017", Hgt: "183cm"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidPassword(tt.pwd)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestPart1Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part1Core(ins)
	assert.EqualValues(t, 216, got)
}
