package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRowNumber(t *testing.T) {
	tests := []struct {
		name string
		bd   string
		want int
	}{
		{"1", "BFFFBBF", 70},
		{"2", "FFFBBBF", 14},
		{"3", "BBFFBBF", 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRowNumber(tt.bd, 0, 127)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestGetColumnNumber(t *testing.T) {
	tests := []struct {
		name string
		bd   string
		want int
	}{
		{"1", "RRR", 7},
		{"2", "RLL", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getColumnNumber(tt.bd, 0, 7)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestGetSeatID(t *testing.T) {
	tests := []struct {
		name string
		bd   string
		want int
	}{
		{"1", "BFFFBBFRRR", 567},
		{"2", "FFFBBBFRRR", 119},
		{"3", "BBFFBBFRLL", 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSeatID(tt.bd)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestPart1Core(t *testing.T) {
	ins := readInput("input.txt")
	got := part1Core(ins)
	assert.EqualValues(t, 878, got)
}
