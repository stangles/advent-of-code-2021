package aoc

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/stangles/advent-of-code-2021/util"
)

type Line struct {
	a, b Point
}

type Point struct {
	x, y int
}

func Day5Part1() int {
	lineInput := util.MustString(util.GetStringInput("data/day5.txt"))

	lines, xMax, yMax := parseLines(lineInput, true)

	grid := make([][]int, yMax+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, xMax+1)
	}

	for _, line := range lines {
		if line.a.x == line.b.x {
			for i := min(line.a.y, line.b.y); i <= max(line.a.y, line.b.y); i++ {
				grid[i][line.a.x] += 1
			}
		} else {
			for i := min(line.a.x, line.b.x); i <= max(line.a.x, line.b.x); i++ {
				grid[line.a.y][i] += 1
			}
		}
	}

	count := 0
	for _, row := range grid {
		for _, col := range row {
			if col >= 2 {
				count++
			}
		}
	}

	return count
}

func Day5Part2() int {
	lineInput := util.MustString(util.GetStringInput("data/day5.txt"))
	lines, xMax, yMax := parseLines(lineInput, false)

	fmt.Println(len(lines), xMax, yMax)
	grid := make([][]int, yMax+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, xMax+1)
	}

	for _, line := range lines {
		if line.a.x == line.b.x {
			for i := min(line.a.y, line.b.y); i <= max(line.a.y, line.b.y); i++ {
				grid[i][line.a.x] += 1
			}
		} else if line.a.y == line.b.y {
			for i := min(line.a.x, line.b.x); i <= max(line.a.x, line.b.x); i++ {
				grid[line.a.y][i] += 1
			}
		} else {
			x := line.a.x
			y := line.a.y
			xAdj := (line.b.x - line.a.x) / int(math.Abs(float64(line.b.x)-float64(line.a.x)))
			yAdj := (line.b.y - line.a.y) / int(math.Abs(float64(line.b.y)-float64(line.a.y)))
			for x != line.b.x && y != line.b.y {
				grid[y][x] += 1
				x += xAdj
				y += yAdj
			}
			grid[line.b.y][line.b.x] += 1
		}
	}

	count := 0
	for _, row := range grid {
		for _, col := range row {
			if col >= 2 {
				count++
			}
		}
	}

	return count
}

func parseLines(lineInput []string, filterNonVertHor bool) ([]Line, int, int) {
	lines := make([]Line, 0)
	xMax := -1
	yMax := -1
	for _, line := range lineInput {
		points := strings.Split(line, " -> ")
		p1 := strings.Split(points[0], ",")
		p2 := strings.Split(points[1], ",")

		l := Line{
			a: Point{
				x: mustInt(strconv.Atoi(p1[0])),
				y: mustInt(strconv.Atoi(p1[1])),
			},
			b: Point{
				x: mustInt(strconv.Atoi(p2[0])),
				y: mustInt(strconv.Atoi(p2[1])),
			},
		}
		if filterNonVertHor {
			if p1[0] == p2[0] || p1[1] == p2[1] {
				xMax = max(xMax, l.a.x, l.b.x)
				yMax = max(yMax, l.a.y, l.b.y)
				lines = append(lines, l)
			}
		} else {
			xMax = max(xMax, l.a.x, l.b.x)
			yMax = max(yMax, l.a.y, l.b.y)
			lines = append(lines, l)
		}
	}

	return lines, xMax, yMax
}

func mustInt(num int, err error) int {
	if err != nil {
		panic(err)
	}
	return num
}

func max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
