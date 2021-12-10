package aoc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/stangles/advent-of-code-2021/util"
)

var easy = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

func Day8Part1() int {
	input := util.MustString(util.GetStringInput("data/day8.txt"))
	outputs := make([]string, 0)

	for _, line := range input {
		outputs = append(outputs, strings.Split(line, " | ")[1])
	}

	count := 0
	for _, output := range outputs {
		for _, digit := range strings.Fields(output) {
			if _, isEasy := easy[len(digit)]; isEasy {
				count++
			}
		}
	}
	return count
}

type segmentPart int

const (
	top segmentPart = 1 << iota
	topLeft
	topRight
	middle
	bottomLeft
	bottomRight
	bottom
)

var segmentsToDigits = map[segmentPart]int{
	top | topLeft | topRight | bottomLeft | bottomRight | bottom: 0,
	topRight | bottomRight:                                                1,
	top | topRight | middle | bottomLeft | bottom:                         2,
	top | topRight | middle | bottomRight | bottom:                        3,
	topLeft | middle | topRight | bottomRight:                             4,
	top | topLeft | middle | bottomRight | bottom:                         5,
	top | topLeft | middle | bottomLeft | bottomRight | bottom:            6,
	top | topRight | bottomRight:                                          7,
	top | topLeft | topRight | middle | bottomLeft | bottomRight | bottom: 8,
	top | topLeft | topRight | middle | bottomRight | bottom:              9,
}

type signal map[string]bool

func Day8Part2() int {
	input := util.MustString(util.GetStringInput("data/day8.txt"))

	outputTotal := 0
	for _, line := range input {
		split := strings.Split(line, " | ")
		signalPatterns := strings.Fields(split[0])
		outputValues := strings.Fields(split[1])
		signals := toSignals(signalPatterns)
		signalMapping := make(map[int]signal)

		// seed signal mapping with known signals
		for _, signal := range signals {
			length := len(signal)
			switch length {
			case 2:
				signalMapping[1] = signal
			case 3:
				signalMapping[7] = signal
			case 4:
				signalMapping[4] = signal
			case 7:
				signalMapping[8] = signal
			}
		}

		for _, signal := range signals {
			if len(signal) == 6 {
				if len(signal.minus(signalMapping[4])) == 2 {
					signalMapping[9] = signal
				} else {
					remainder := len(signal.minus(signalMapping[7]))
					if remainder == 3 {
						signalMapping[0] = signal
					} else if remainder == 4 {
						signalMapping[6] = signal
					}
				}
			}
		}

		configuration := make(map[segmentPart]signal)
		configuration[top] = signalMapping[7].minus(signalMapping[1])
		configuration[topRight] = signalMapping[1]
		configuration[bottomRight] = signalMapping[1]
		configuration[topLeft] = signalMapping[4].minus(signalMapping[1])
		configuration[middle] = signalMapping[4].minus(signalMapping[1])
		configuration[bottomLeft] = signalMapping[8].minus(signalMapping[4]).minus(configuration[top])
		configuration[bottom] = signalMapping[8].minus(signalMapping[4]).minus(configuration[top])

		tmpBottom := signalMapping[9].minus(signalMapping[4]).minus(signalMapping[7])
		configuration[bottomLeft] = configuration[bottomLeft].minus(tmpBottom)
		configuration[bottom] = tmpBottom

		tmpMiddle := signalMapping[8].minus(signalMapping[0])
		configuration[topLeft] = configuration[topLeft].minus(tmpMiddle)
		configuration[middle] = tmpMiddle

		for _, signal := range signals {
			if len(signal) == 5 {
				remainder := signal.minus(signalMapping[6])
				if len(remainder) != 0 {
					configuration[topRight] = remainder
					configuration[bottomRight] = configuration[bottomRight].minus(configuration[topRight])
				}
			}
		}

		strOutput := ""
		for _, outputValue := range outputValues {
			var segmentValue segmentPart
			for i := 0; i < len(outputValue); i++ {
				for segment, signal := range configuration {
					if signal[string(outputValue[i])] {
						segmentValue |= segment
					}
				}
			}
			strOutput += fmt.Sprint(segmentsToDigits[segmentValue])
		}

		total, err := strconv.Atoi(strOutput)
		if err != nil {
			log.Fatalf("unable to convert total %s: %v", strOutput, err)
		}
		outputTotal += total
	}
	return outputTotal
}

func toSignals(signalPatterns []string) []signal {
	signals := make([]signal, 0)
	for _, signalPattern := range signalPatterns {
		signal := make(signal)
		for i := 0; i < len(signalPattern); i++ {
			signal[string(signalPattern[i])] = true
		}
		signals = append(signals, signal)
	}

	return signals
}

func (s signal) minus(o signal) signal {
	result := make(signal)
	for sign := range s {
		if _, contains := o[sign]; !contains {
			result[sign] = true
		}
	}

	return result
}
