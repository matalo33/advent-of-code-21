package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var reports []string

	for scanner.Scan() {
		reports = append(reports, scanner.Text())
	}

	fmt.Printf("Power consumption: %v\n", powerConsumption(reports))
	fmt.Printf("Life support rating: %v\n", (findRating(reports, byte('1'), byte('0'), true) *
		findRating(reports, byte('0'), byte('1'), false)))
}

func powerConsumption(reports []string) int64 {
	var gamma, epsilon []rune
	for i := 0; i < len(reports[0]); i++ {
		bits := 0
		for j := 0; j < len(reports); j++ {
			if reports[j][i] == '1' {
				bits += 1
			}
		}
		if bits > len(reports)/2 {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		} else {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		}
	}

	g, _ := strconv.ParseInt(string(gamma), 2, 0)
	e, _ := strconv.ParseInt(string(epsilon), 2, 0)

	return g * e
}

func findRating(reports []string, def byte, alt byte, moreOrLess bool) int64 {
	oxyReports := reports
	for i := 0; i < len(reports[0]); i++ {
		zeros, ones := 0, 0
		for j := 0; j < len(oxyReports); j++ {
			if oxyReports[j][i] == def {
				ones += 1
			} else {
				zeros += 1
			}
		}
		filter := byte(def)
		if moreOrLess {
			if zeros > ones {
				filter = byte(alt)
			}
		} else {
			if zeros < ones {
				filter = byte(alt)
			}
		}
		var newOxyreports []string
		for j := 0; j < len(oxyReports); j++ {
			if oxyReports[j][i] == filter {
				newOxyreports = append(newOxyreports, oxyReports[j])
			}
		}
		if len(newOxyreports) == 1 {
			result, _ := strconv.ParseInt(string(newOxyreports[0]), 2, 0)
			return result
		}
		oxyReports = newOxyreports
	}

	return 0
}
