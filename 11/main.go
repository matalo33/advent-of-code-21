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

	octopi := make([][]int, 10)
	rowN := 0
	for scanner.Scan() {
		row := make([]int, 10)
		for k, v := range scanner.Text() {
			i, _ := strconv.Atoi(string(v))
			row[k] = i
		}
		octopi[rowN] = row
		rowN++
	}

	fmt.Printf("Number of flashes: %v\n", countFlashes(octopi, 100))
}

func countFlashes(octopi [][]int, steps int) int {
	for step := 0; step < steps; step++ {
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				octopi[x][y]++
			}
		}
	}
	return 0
}
