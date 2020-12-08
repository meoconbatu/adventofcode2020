package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part2Core(ins)
	assert.EqualValues(t, 1250, got)
}
