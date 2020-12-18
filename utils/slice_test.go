package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		name   string
		s1, s2 []int
		want   []int
	}{
		{"1", []int{1, 3}, []int{1, 3}, []int{1, 3}},
		{"2", []int{1, 3}, []int{2, 3}, []int{3}},
		{"3", []int{1, 3}, []int{2, 4}, []int{}},
		{"4", []int{1, 3, 4}, []int{2, 4}, []int{4}},
		{"5", []int{3, 4, 6}, []int{2, 4, 5, 6}, []int{4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindIntersection(tt.s1, tt.s2)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

// func TestPart2Core(t *testing.T) {
// 	ins := readInput("input.txt")
// 	got := part2Core(ins)
// 	assert.EqualValues(t, 441, got)
// }
