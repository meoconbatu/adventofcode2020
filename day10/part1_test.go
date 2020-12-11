package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part1Core(ins)
	assert.EqualValues(t, 2112, got)
}