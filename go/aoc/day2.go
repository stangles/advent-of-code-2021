package aoc

import (
	"log"
	"strconv"
	"strings"

	"github.com/stangles/advent-of-code-2021/util"
)

func Day2Part1() int {
	commands, err := util.GetStringInput("data/day2.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	horizontal := 0
	vertical := 0
	for _, command := range commands {
		parsed := strings.Split(command, " ")
		dir := parsed[0]
		distance, err := strconv.Atoi(parsed[1])
		if err != nil {
			log.Fatalf("unable to convert string to int: %v", err)
		}

		if dir == "forward" {
			horizontal += distance
		} else if dir == "up" {
			vertical -= distance
		} else if dir == "down" {
			vertical += distance
		}
	}

	return horizontal * vertical
}

func Day2Part2() int {
	commands, err := util.GetStringInput("data/day2.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	horizontal := 0
	depth := 0
	aim := 0
	for _, command := range commands {
		parsed := strings.Split(command, " ")
		dir := parsed[0]
		distance, err := strconv.Atoi(parsed[1])
		if err != nil {
			log.Fatalf("unable to convert string to int: %v", err)
		}

		if dir == "forward" {
			horizontal += distance
			depth += aim * distance
		} else if dir == "up" {
			aim -= distance
		} else if dir == "down" {
			aim += distance
		}
	}

	return horizontal * depth
}
