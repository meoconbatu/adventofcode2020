package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part1Core(ins)
	assert.EqualValues(t, 1797, got)
}
