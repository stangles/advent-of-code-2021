package aoc

import (
	"github.com/stangles/advent-of-code-2021/util"
)

const boardSize = 5

func Day4Part1() int {
	nums := util.MustInt(util.GetIntInputWithSplitFunc("data/day4Numbers.csv", util.ScanCommaSeparated))
	boards := util.MustBoard(util.GetBoards("data/day4Boards.txt", boardSize))

	minSet := make(map[int]bool)
	for i := 0; i < boardSize; i++ {
		minSet[nums[i]] = true
	}

	for i := boardSize - 1; i < len(nums); i++ {
		minSet[nums[i]] = true
		for _, board := range boards {
			if checkBoard(board, minSet) {
				return nums[i] * sumUnmarked(board, minSet)
			}
		}
	}

	panic("input did not contain winning board")
}

func Day4Part2() int {
	nums := util.MustInt(util.GetIntInputWithSplitFunc("data/day4Numbers.csv", util.ScanCommaSeparated))
	boards := util.MustBoard(util.GetBoards("data/day4Boards.txt", boardSize))

	minSet := make(map[int]bool)
	for i := 0; i < boardSize; i++ {
		minSet[nums[i]] = true
	}

	winners := make(map[int]int, 0)
	lastEntry := -1
	for i := boardSize - 1; i < len(nums); i++ {
		minSet[nums[i]] = true
		for pos, board := range boards {
			if _, isWinner := winners[pos]; !isWinner && checkBoard(board, minSet) {
				winners[pos] = nums[i]
				lastEntry = pos
			}
		}

		if len(winners) >= len(boards) {
			break
		}
	}

	return winners[lastEntry] * sumUnmarked(boards[lastEntry], minSet)
}

func checkBoard(board [][]int, nums map[int]bool) bool {
	for _, row := range board {
		matched := true
		for _, col := range row {
			if _, ok := nums[col]; !ok {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}

	for col := 0; col < boardSize; col++ {
		matched := true
		for row := 0; row < boardSize; row++ {
			if _, ok := nums[board[row][col]]; !ok {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}
	return false
}

func sumUnmarked(board [][]int, called map[int]bool) int {
	sum := 0
	for _, row := range board {
		for _, col := range row {
			if _, isNumCalled := called[col]; !isNumCalled {
				sum += col
			}
		}
	}

	return sum
}
