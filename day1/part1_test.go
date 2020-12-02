package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part1Core(ins, 2020)
	assert.EqualValues(t, 468051, got)
}
