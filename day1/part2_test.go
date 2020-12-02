package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part2Core(ins, 2020)
	assert.EqualValues(t, 272611658, got)
}
