package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountTree(t *testing.T) {
	tests := []struct {
		name string
		ins  [][]int
		want int
	}{
		{"1", [][]int{{0, 0, 0, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTree(tt.ins, 3, 1)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestPart1Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part1Core(ins)
	assert.EqualValues(t, 234, got)
}
