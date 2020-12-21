package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc2(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       int
	}{
		{"1", "3 * 6 + 8 + 6 + (5 * 8 + 8)", 300},
		{"2", "(2 * 2) + 5", 9},
		{"3", "((2 * 2) + 5)", 9},
		{"4", "((2 * 2 + 1) + 5)", 11},
		{"5", "2 + 2 * 5", 20},
		{"5", "(2 + (2 * 5))", 12},
		{"5", "1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"5", "2 * 3 + (4 * 5)", 46},
		{"5", "5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5", "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"5", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc2(tt.expression)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestPart2Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part2Core(ins)
	assert.EqualValues(t, 461295257566346, got)
}
