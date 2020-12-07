package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAnswerYes2(t *testing.T) {
	tests := []struct {
		name  string
		group []string
		want  int
	}{
		{"1", []string{"abc"}, 3},
		{"2", []string{"a", "b", "c"}, 0},
		{"3", []string{"ab", "ac"}, 1},
		{"4", []string{"a", "a", "a", "a"}, 1},
		{"5", []string{"b"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countAnswerYes2(tt.group)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestPart2Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part2Core(ins)
	assert.EqualValues(t, 3437, got)
}
