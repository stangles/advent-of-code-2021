package aoc

import (
	"math"

	"github.com/stangles/advent-of-code-2021/util"
)

func Day7Part1() int {
	return day7(func(distance int) int {
		return distance
	})
}

func Day7Part2() int {
	return day7(func(distance int) int {
		return distance * (distance + 1) / 2
	})
}

func day7(computeFunc func(int) int) int {
	crabs := util.MustInt(util.GetIntInputWithSplitFunc("data/day7.csv", util.ScanCommaSeparated))

	min := util.Min(crabs...)
	max := util.Max(crabs...)

	minFuel := math.MaxInt64
	for pos := min; pos <= max; pos++ {
		fuel := 0
		for _, crab := range crabs {
			distance := int(math.Abs(float64(pos - crab)))
			fuel += computeFunc(distance)
		}
		minFuel = util.Min(fuel, minFuel)
	}

	return minFuel
}
