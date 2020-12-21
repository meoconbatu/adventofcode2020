package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       int
	}{
		{"1", "1 + 2 + 3", 6},
		{"2", "1 + (2 + 3)", 6},
		{"3", "(1 + (2 + 3))", 6},
		{"4", "(1 + 2 * (2 + 3))", 15},
		{"5", "1+5", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc(tt.expression)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestPart1Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part1Core(ins)
	assert.EqualValues(t, 15285807527593, got)
}
