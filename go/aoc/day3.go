package aoc

import (
	"math"

	"github.com/stangles/advent-of-code-2021/util"
)

const bitCount = 12

func Day3Part1() int {
	binaryStrings := util.MustString(util.GetStringInput("data/day3.txt"))
	nums := util.MustBinaryStrToUint16(util.BinaryStrToUint16(binaryStrings))
	onesCounts := getOnesCounts(nums)

	gamma := 0xFFF
	epsilon := 0xFFF
	for pos, count := range onesCounts {
		if count > len(nums)/2 { // 1 is the most common in this pos, therefore 0 is the least common
			epsilon &= ^int(math.Pow(2, float64(pos)))
		} else { // 0 is most common, therefore 1 is least common
			gamma &= ^int(math.Pow(2, float64(pos)))
		} // we're not really considering the case of a tie I guess :)
	}
	return gamma * epsilon
}

func Day3Part2() int {
	binaryStrings := util.MustString(util.GetStringInput("data/day3.txt"))
	nums := util.MustBinaryStrToUint16(util.BinaryStrToUint16(binaryStrings))

	return int(getRating(nums, true)) * int(getRating(nums, false))
}

func getOnesCounts(nums []uint16) []int {
	onesCounts := make([]int, bitCount)
	for _, num := range nums {
		for pos := len(onesCounts) - 1; pos >= 0; pos-- {
			mask := uint16(math.Pow(2, float64(pos)))
			onesCounts[pos] += int(num & mask >> pos)
		}
	}

	return onesCounts
}

func getOnesCountAt(idx int, nums []uint16) int {
	return getOnesCounts(nums)[idx]
}

func getRating(nums []uint16, o2 bool) uint16 {
	ratingCandidates := make([]uint16, 0)
	for _, num := range nums {
		ratingCandidates = append(ratingCandidates, num)
	}

	for idx := bitCount - 1; idx >= 0; idx-- {
		countOnes := getOnesCountAt(idx, ratingCandidates)
		countZeros := len(ratingCandidates) - countOnes

		newCandidates := make([]uint16, 0)
		for _, num := range ratingCandidates {
			mask := uint16(math.Pow(2, float64(idx)))
			bitSet := int((num&mask)>>uint(idx)) == 1
			if o2 {
				if countOnes >= countZeros && bitSet {
					newCandidates = append(newCandidates, num)
				} else if countZeros > countOnes && !bitSet {
					newCandidates = append(newCandidates, num)
				}
			} else {
				if countZeros <= countOnes && !bitSet {
					newCandidates = append(newCandidates, num)
				} else if countOnes < countZeros && bitSet {
					newCandidates = append(newCandidates, num)
				}
			}
		}

		ratingCandidates = newCandidates
		if len(ratingCandidates) == 1 {
			break
		}
	}

	return ratingCandidates[0]
}
