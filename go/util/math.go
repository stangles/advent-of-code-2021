package util

import "math"

func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func Min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
