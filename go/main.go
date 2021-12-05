package main

import (
	"fmt"
	"time"

	"github.com/stangles/advent-of-code-2021/aoc"
)

func main() {
	ret := withTimings(func() interface{} {
		return aoc.Day5Part2()
	})
	fmt.Println(ret)
}

func withTimings(f func() interface{}) interface{} {
	start := time.Now()
	ret := f()
	fmt.Println("done in", time.Now().Sub(start))
	return ret
}
