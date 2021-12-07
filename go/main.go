package main

import (
	"fmt"

	"github.com/stangles/advent-of-code-2021/aoc"
	"github.com/stangles/advent-of-code-2021/util"
)

func main() {
	ret := util.WithTimings(func() interface{} {
		return aoc.Day7Part1()
	})
	fmt.Println(ret)
}
