package aoc

import (
	"log"

	"github.com/stangles/advent-of-code-2021/util"
)

func Day1Part1() int {
	lines, err := util.GetIntInput("data/day1.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	prev := lines[0]
	count := 0
	for i := 1; i < len(lines); i++ {
		if lines[i] > prev {
			count++
		}
		prev = lines[i]
	}

	return count
}

func Day1Part2() int {
	lines, err := util.GetIntInput("data/day1.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	count := 0
	sum := lines[0] + lines[1] + lines[2]
	for i := 3; i < len(lines); i++ {
		newSum := sum + lines[i] - lines[i-3]
		if newSum > sum {
			count++
		}
		sum = newSum
	}

	return count
}
