// A lot of inspriation taken from https://github.com/nlowe/aoc2021/blob/master/challenge/day8/b.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) (count int) {
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		outputDigits := strings.Split(parts[1], " ")
		for _, output := range outputDigits {
			if len(output) == 2 ||
				len(output) == 4 ||
				len(output) == 3 ||
				len(output) == 7 {
				count++
			}
		}
	}
	return count
}

func part2(lines []string) (total int) {
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		signals, outputs := sortSlice(strings.Split(parts[0], " ")), sortSlice(strings.Split(parts[1], " "))
		total += decode(signals, outputs)
	}
	return total
}

func decode(signals []string, outputs []string) (output int) {
	coding := map[int]string{}

	// 1, 4, 7 and 8 are unique
	for _, signal := range signals {
		switch len(signal) {
		case 2:
			coding[1] = signal
		case 3:
			coding[7] = signal
		case 4:
			coding[4] = signal
		case 7:
			coding[8] = signal
		}
	}

	// Find 3, by looking at 5chars and matching 1's signals
	for _, signal := range signals {
		if len(signal) == 5 {
			if len(leftNotInRight(coding[1], signal)) == 0 {
				coding[3] = signal
			}
		}
	}

	for _, signal := range signals {
		if len(signal) == 6 {
			// Find 6, 1 of 1 is not shared
			if len(leftNotInRight(coding[1], signal)) == 1 {
				coding[6] = signal
			}
			// Find 9, 0 of 4 are shared
			if len(leftNotInRight(coding[4], signal)) == 0 {
				coding[9] = signal
			}
			// Find 0, 1 of 4 is not shared
			if len(leftNotInRight(coding[4], signal)) == 1 {
				coding[0] = signal
			}
		}
	}

	// Determine top right by diffing between 1 and 6
	segmentC := leftNotInRight(coding[1], coding[6])

	for _, signal := range signals {
		if len(signal) == 5 {
			// already have 3
			if signal != coding[3] {
				if strings.Contains(signal, segmentC) {
					// Top right lit means this is a 2
					coding[2] = signal
				} else {
					// Otherwise its a 5
					coding[5] = signal
				}
			}
		}
	}

	for _, outDigit := range outputs {
		output *= 10
		for k, v := range coding {
			if v == outDigit {
				output += k
			}
		}
	}
	return output
}

func leftNotInRight(left, right string) (missing string) {
loop:
	for _, l := range left {
		for _, r := range right {
			if l == r {
				continue loop
			}
		}
		//count++
		missing += string(l)
	}
	return missing
}

func sortSlice(slice []string) []string {
	result := make([]string, len(slice))

	for i, str := range slice {
		v := []rune(str)
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})

		result[i] = string(v)
	}

	return result
}
