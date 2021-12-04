package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetIntInput(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file '%s': %w", filename, err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file '%s': %v", filename, err)
		}
	}()

	lines := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		next, err := strconv.Atoi(text)
		if err != nil {
			return nil, fmt.Errorf("unable to convert value %v to int: %w", text, err)
		}

		lines = append(lines, next)
	}

	return lines, nil
}

func GetStringInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file '%s': %w", filename, err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file '%s': %v", filename, err)
		}
	}()

	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("unable to read contents of file '%s': %w", filename, err)
	}

	lines := strings.Split(string(contents), "\n")
	if lines[len(lines)-1] == "" {
		return lines[:len(lines)-1], nil
	} else {
		return lines, nil
	}
}

func MustInt(ints []int, err error) []int {
	if err != nil {
		panic(err)
	}
	return ints
}

func MustString(strings []string, err error) []string {
	if err != nil {
		panic(err)
	}
	return strings
}

func BinaryStrToUint16(binaryStrings []string) ([]uint16, error) {
	nums := make([]uint16, 0)
	for _, bin := range binaryStrings {
		num, err := strconv.ParseUint(bin, 2, 16)
		if err != nil {
			return nil, fmt.Errorf("unable to convert %s to number: %w", bin, err)
		}

		nums = append(nums, uint16(num))
	}

	return nums, nil
}

func MustBinaryStrToUint16(nums []uint16, err error) []uint16 {
	if err != nil {
		panic(err)
	}
	return nums
}
