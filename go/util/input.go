package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetIntInput(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file '%s': %w", filename, err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
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
