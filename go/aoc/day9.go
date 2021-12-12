package aoc

import (
	"container/heap"
	"log"
	"strconv"

	"github.com/stangles/advent-of-code-2021/util"
)

func Day9Part1() int {
	heightmapStr := util.MustString(util.GetStringInput("data/day9.txt"))
	heightmap := strToInt(heightmapStr)

	riskLevel := 0
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			cur := heightmap[i][j]
			if j > 0 {
				if heightmap[i][j-1] <= cur {
					continue
				}
			}
			if j < len(heightmap[i])-1 {
				if heightmap[i][j+1] <= cur {
					continue
				}
			}

			if i > 0 {
				if heightmap[i-1][j] <= cur {
					continue
				}
			}
			if i < len(heightmap)-1 {
				if heightmap[i+1][j] <= cur {
					continue
				}
			}
			riskLevel += cur + 1
		}
	}
	return riskLevel
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type basinNeighbor struct {
	x, y int
}

func Day9Part2() int {
	heightmapStr := util.MustString(util.GetStringInput("data/day9.txt"))
	heightmap := strToInt(heightmapStr)

	topThree := &intHeap{}
	heap.Init(topThree)
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			cur := heightmap[i][j]
			if j > 0 {
				if heightmap[i][j-1] <= cur {
					continue
				}
			}
			if j < len(heightmap[i])-1 {
				if heightmap[i][j+1] <= cur {
					continue
				}
			}
			if i > 0 {
				if heightmap[i-1][j] <= cur {
					continue
				}
			}
			if i < len(heightmap)-1 {
				if heightmap[i+1][j] <= cur {
					continue
				}
			}

			count := 1
			seen := make(map[basinNeighbor]bool, 0)
			neighbors := getBasinNeighbors(heightmap, j, i)
			for idx := 0; idx < len(neighbors); idx++ {
				if _, isKnown := seen[neighbors[idx]]; !isKnown {
					seen[neighbors[idx]] = true
					count++
					neighbors = append(neighbors, getBasinNeighbors(heightmap, neighbors[idx].x, neighbors[idx].y)...)
				}
			}
			heap.Push(topThree, count)
		}
	}

	return heap.Pop(topThree).(int) * heap.Pop(topThree).(int) * heap.Pop(topThree).(int)
}

func getBasinNeighbors(heightmap [][]int, x, y int) []basinNeighbor {
	neighbors := make([]basinNeighbor, 0)
	cur := heightmap[y][x]
	if y > 0 {
		neighbor := heightmap[y-1][x]
		if neighbor > cur && neighbor < 9 {
			neighbors = append(neighbors, basinNeighbor{x: x, y: y - 1})
		}
	}
	if y < len(heightmap)-1 {
		neighbor := heightmap[y+1][x]
		if neighbor > cur && neighbor < 9 {
			neighbors = append(neighbors, basinNeighbor{x: x, y: y + 1})
		}
	}

	if x > 0 {
		neighbor := heightmap[y][x-1]
		if neighbor > cur && neighbor < 9 {
			neighbors = append(neighbors, basinNeighbor{x: x - 1, y: y})
		}
	}
	if x < len(heightmap[y])-1 {
		neighbor := heightmap[y][x+1]
		if neighbor > cur && neighbor < 9 {
			neighbors = append(neighbors, basinNeighbor{x: x + 1, y: y})
		}
	}

	return neighbors
}

func strToInt(heightmapStr []string) [][]int {
	heightmap := make([][]int, 0)
	for i, row := range heightmapStr {
		heightmap = append(heightmap, make([]int, 0))
		for j := 0; j < len(row); j++ {
			point, err := strconv.Atoi(string(row[j]))
			if err != nil {
				log.Fatalf("failed to convert %s to numeric type: %v", string(row[j]), err)
			}
			heightmap[i] = append(heightmap[i], point)
		}
	}

	return heightmap
}
