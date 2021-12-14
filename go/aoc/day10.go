package aoc

import (
	"sort"

	"github.com/stangles/advent-of-code-2021/util"
)

var points = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var completionPoints = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

var openTokens = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var opposites = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

func Day10Part1() int {
	commands := util.MustString(util.GetStringInput("data/day10.txt"))
	score := 0
	for _, command := range commands {
		opens := make([]int, 0)
		for i := 0; i < len(command); i++ {
			cur := string(command[i])
			if _, isOpen := openTokens[cur]; isOpen {
				opens = append(opens, i)
			} else {
				if len(opens) > 0 && string(command[opens[len(opens)-1]]) == opposites[cur] {
					opens = opens[0 : len(opens)-1]
				} else {
					score += points[cur]
					break
				}
			}
		}
	}
	return score
}

func Day10Part2() int {
	commands := util.MustString(util.GetStringInput("data/day10.txt"))
	scores := make([]int, 0)
	for _, command := range commands {
		opens := make([]int, 0)
		for i := 0; i < len(command); i++ {
			cur := string(command[i])
			if _, isOpen := openTokens[cur]; isOpen {
				opens = append(opens, i)
			} else {
				if len(opens) > 0 && string(command[opens[len(opens)-1]]) == opposites[cur] {
					opens = opens[0 : len(opens)-1]
				} else {
					opens = make([]int, 0)
					break
				}
			}
		}

		if len(opens) > 0 {
			completion := ""
			for i := len(opens) - 1; i >= 0; i-- {
				completion += openTokens[string(command[opens[i]])]
			}
			score := 0
			for i := 0; i < len(completion); i++ {
				score *= 5
				score += completionPoints[string(completion[i])]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)

	return scores[len(scores)/2]
}
