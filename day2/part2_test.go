package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountValidPassword2(t *testing.T) {
	tests := []struct {
		name     string
		policies []passwordPolicy
		want     int
	}{
		{"1", []passwordPolicy{{1, 3, 'a', "abcde"}, {1, 3, 'b', "cdefg"}, {2, 9, 'c', "ccccccccc"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countValidPassword2(tt.policies)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestPart2Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part2Core(ins)
	assert.EqualValues(t, 441, got)
}
