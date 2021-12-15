package aoc

import (
	"log"
	"strconv"

	"github.com/stangles/advent-of-code-2021/util"
)

func Day11Part1() int {
	octopuses := util.MustString(util.GetStringInput("data/day11.txt"))
	grid := make([][]int, 0)

	for _, octoLine := range octopuses {
		octoInts := make([]int, len(octoLine))
		for i := 0; i < len(octoLine); i++ {
			octopus, err := strconv.Atoi(string(octoLine[i]))
			if err != nil {
				log.Fatalf("unable to convert %s: %v", string(octoLine[i]), err)
			}

			octoInts[i] = octopus
		}
		grid = append(grid, octoInts)
	}

	count := 0
	for iterations := 0; iterations < 1000; iterations++ {
		increment(grid)
		count += flash(grid)
	}

	return count
}

func Day11Part2() int {
	octopuses := util.MustString(util.GetStringInput("data/day11.txt"))
	grid := make([][]int, 0)

	for _, octoLine := range octopuses {
		octoInts := make([]int, len(octoLine))
		for i := 0; i < len(octoLine); i++ {
			octopus, err := strconv.Atoi(string(octoLine[i]))
			if err != nil {
				log.Fatalf("unable to convert %s: %v", string(octoLine[i]), err)
			}

			octoInts[i] = octopus
		}
		grid = append(grid, octoInts)
	}

	step := 1
	for {
		increment(grid)
		iterCount := flash(grid)
		if iterCount == len(grid)*len(grid[0]) {
			return step
		}
		step++
	}
}

func increment(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j]++
		}
	}
}

func flash(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 9 {
				count += flashOne(grid, i, j)
			}
		}
	}
	return count
}

func flashOne(grid [][]int, i, j int) int {
	grid[i][j] = 0
	count := 1
	if i > 0 {
		if grid[i-1][j] != 0 {
			grid[i-1][j]++
		}
		if grid[i-1][j] > 9 {
			count += flashOne(grid, i-1, j)
		}
		if j > 0 {
			if grid[i-1][j-1] != 0 {
				grid[i-1][j-1]++
			}
			if grid[i-1][j-1] > 9 {
				count += flashOne(grid, i-1, j-1)
			}
		}
		if j < len(grid[i])-1 {
			if grid[i-1][j+1] != 0 {
				grid[i-1][j+1]++
			}
			if grid[i-1][j+1] > 9 {
				count += flashOne(grid, i-1, j+1)
			}
		}
	}
	if j > 0 {
		if grid[i][j-1] != 0 {
			grid[i][j-1]++
		}
		if grid[i][j-1] > 9 {
			count += flashOne(grid, i, j-1)
		}
	}
	if j < len(grid[i])-1 {
		if grid[i][j+1] != 0 {
			grid[i][j+1]++
		}
		if grid[i][j+1] > 9 {
			count += flashOne(grid, i, j+1)
		}
	}
	if i < len(grid)-1 {
		if grid[i+1][j] != 0 {
			grid[i+1][j]++
		}
		if grid[i+1][j] > 9 {
			count += flashOne(grid, i+1, j)
		}
		if j > 0 {
			if grid[i+1][j-1] != 0 {
				grid[i+1][j-1]++
			}
			if grid[i+1][j-1] > 9 {
				count += flashOne(grid, i+1, j-1)
			}
		}
		if j < len(grid[i])-1 {
			if grid[i+1][j+1] != 0 {
				grid[i+1][j+1]++
			}
			if grid[i+1][j+1] > 9 {
				count += flashOne(grid, i+1, j+1)
			}
		}
	}

	return count
}
