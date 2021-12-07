package aoc

import (
	"github.com/stangles/advent-of-code-2021/util"
)

func Day6Part1() int {
	return day6(80)
}

func Day6Part2() int {
	return day6(256)
}

func day6(days int) int {
	ages := util.MustInt(util.GetIntInputWithSplitFunc("data/day6.csv", util.ScanCommaSeparated))

	fishies := make(map[int]int, 0)
	for _, age := range ages {
		if _, ok := fishies[age]; !ok {
			fishies[age] = 1
		} else {
			fishies[age] += 1
		}
	}

	next := fishies[0]
	for day := 0; day < days; day++ {
		for age := 8; age >= 0; age-- {
			tmp := fishies[age]
			fishies[age] = next
			next = tmp
			if age == 0 {
				fishies[6] += next
				fishies[8] = next
			}
		}
	}

	sum := 0
	for _, ageCount := range fishies {
		sum += ageCount
	}

	return sum
}
