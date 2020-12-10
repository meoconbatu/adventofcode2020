package utils

import "math"

// FindMinMax func return min and max number of ins array
func FindMinMax(ins []int) (int, int) {
	min, max := math.MaxInt64, math.MinInt64
	for i := 0; i < len(ins); i++ {
		if ins[i] > max {
			max = ins[i]
		}
		if ins[i] < min {
			min = ins[i]
		}
	}
	return min, max
}
