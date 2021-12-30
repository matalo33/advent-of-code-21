package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type coord struct {
	y int
	x int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var heightmap [][]int
	for scanner.Scan() {
		row := make([]int, len(scanner.Text()))
		for k, v := range scanner.Text() {
			i, _ := strconv.Atoi(string(v))
			row[k] = i
		}
		heightmap = append(heightmap, row)
	}

	lowPoints := getLowPoints(heightmap)

	fmt.Printf("Risk level of lowest points: %v\n", getRiskLevel(lowPoints, heightmap))
	fmt.Printf("Risk level of basins: %v\n", getBasinsRiskLevel(lowPoints, heightmap))
}

func getLowPoints(heightmap [][]int) []coord {
	var lowPoints []coord
	for y := range heightmap {
		for x := range heightmap[y] {
			valid := true
			for {
				// y-1, x
				if y-1 >= 0 {
					if heightmap[y-1][x] <= heightmap[y][x] {
						valid = false
						break
					}
				}
				// y+1, x
				if y+1 <= len(heightmap)-1 {
					if heightmap[y+1][x] <= heightmap[y][x] {
						valid = false
						break
					}
				}
				// y, x-1
				if x-1 >= 0 {
					if heightmap[y][x-1] <= heightmap[y][x] {
						valid = false
						break
					}
				}
				// y, x+1
				if x+1 <= len(heightmap[y])-1 {
					if heightmap[y][x+1] <= heightmap[y][x] {
						valid = false
						break
					}
				}
				break
			}
			if valid {
				lowPoints = append(lowPoints, coord{y, x})
			}
		}
	}
	return lowPoints
}

func getRiskLevel(coords []coord, heightmap [][]int) (res int) {
	for _, coord := range coords {
		res += heightmap[coord.y][coord.x] + 1
	}
	return
}

func getBasinsRiskLevel(coords []coord, heightmap [][]int) (res int) {
	var basinSizes []int
	for _, lowPoint := range coords {
		// a map of seen coords, starting with the lowpoint
		seen := map[coord]struct{}{lowPoint: {}}
		basinSize := 1

		coordsToCheck := []coord{
			{x: lowPoint.x, y: lowPoint.y - 1}, // y-1, x
			{x: lowPoint.x, y: lowPoint.y + 1}, // y+1, x
			{x: lowPoint.x - 1, y: lowPoint.y}, // y, x-1
			{x: lowPoint.x + 1, y: lowPoint.y}, // y, x+1
		}

		// for gets re-evaluated each loop
		for len(coordsToCheck) > 0 {
			// take the current coord off the list
			point := coordsToCheck[0]
			coordsToCheck = coordsToCheck[1:]

			// one liner check for key in map
			if _, ok := seen[point]; ok {
				continue
			}

			// check if out of bounds
			if point.x < 0 ||
				point.x >= len(heightmap[0]) ||
				point.y < 0 ||
				point.y >= len(heightmap) {
				continue
			}

			// height must be below 9
			if heightmap[point.y][point.x] == 9 {
				continue
			}

			// part of the basin!
			basinSize++
			seen[point] = struct{}{}

			// check the neighbours
			coordsToCheck = append(coordsToCheck, []coord{
				{x: point.x, y: point.y - 1}, // y-1, x
				{x: point.x, y: point.y + 1}, // y+1, x
				{x: point.x - 1, y: point.y}, // y, x-1
				{x: point.x + 1, y: point.y}, // y, x+1
			}...)
		}
		basinSizes = append(basinSizes, basinSize)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}
