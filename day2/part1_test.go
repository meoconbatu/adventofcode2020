package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		policies []passwordPolicy
		want     int
	}{
		{"1", []passwordPolicy{{1, 3, 'a', "abcde"}, {1, 3, 'b', "cdefg"}, {2, 9, 'c', "ccccccccc"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countValidPassword(tt.policies)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestPart1Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part1Core(ins)
	assert.EqualValues(t, 467, got)
}
