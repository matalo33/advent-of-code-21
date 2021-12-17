package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var lines []line

	for scanner.Scan() {
		ends := strings.Split(scanner.Text(), " -> ")
		x1, _ := strconv.Atoi(strings.Split(ends[0], ",")[0])
		y1, _ := strconv.Atoi(strings.Split(ends[0], ",")[1])
		x2, _ := strconv.Atoi(strings.Split(ends[1], ",")[0])
		y2, _ := strconv.Atoi(strings.Split(ends[1], ",")[1])
		lines = append(lines, line{x1, y1, x2, y2})
	}

	fmt.Printf("Number of overlapping points: %v\n", overlappingPoints(lines, 1000))
}

func overlappingPoints(lines []line, gridSize int) int {
	grid := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}

	for _, line := range lines {
		if isLineHorizontalOrVertical(line) {
			if line.x1 != line.x2 {
				//horizontal
				min, max := getMinMax(line.x1, line.x2)
				for i := min; i <= max; i++ {
					grid[line.y1][i]++
				}
			} else if line.y1 != line.x2 {
				//verical
				min, max := getMinMax(line.y1, line.y2)
				for i := min; i <= max; i++ {
					grid[i][line.x1]++
				}
			}
		}
	}

	counter := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] >= 2 {
				counter++
			}
		}
	}

	return counter
}

func getMinMax(i1, i2 int) (int, int) {
	if i1 < i2 {
		return i1, i2
	} else {
		return i2, i1
	}
}

func isLineHorizontalOrVertical(line line) bool {
	if line.x1 == line.x2 || line.y1 == line.y2 {
		return true
	}
	return false
}
