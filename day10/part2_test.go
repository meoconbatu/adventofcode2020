package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part2Core(ins)
	assert.EqualValues(t, 3022415986688, got)
}
