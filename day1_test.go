package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1Core(t *testing.T) {
	ins := day1ReadInput("day1_input.txt")
	got := day1Part1Core(ins, 2020)
	assert.EqualValues(t, 468051, got)
}

func TestDay1Part2Core(t *testing.T) {
	ins := day1ReadInput("day1_input.txt")
	got := day1Part2Core(ins, 2020)
	assert.EqualValues(t, 272611658, got)
}
