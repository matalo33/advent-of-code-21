package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	line := strings.Split(string(file), "\n")
	input := strings.Split(line[0], ",")

	fish := make([]int, len(input))
	for i, f := range input {
		fish[i], _ = strconv.Atoi(f)
	}

	fmt.Printf("Number of fish after 80 days: %v\n", makeFishV2(fish, 80))
	fmt.Printf("Number of fish after 256 days: %v\n", makeFishV2(fish, 256))
}

func makeFishV2(fish []int, days int) int {
	cycle := make([]int, 9)
	for _, f := range fish {
		cycle[f]++
	}

	for day := 0; day < days; day++ {
		newFish := cycle[0]
		for i := 1; i <= 8; i++ {
			cycle[i-1] = cycle[i]
		}
		cycle[6] += newFish
		cycle[8] = newFish
	}

	result := 0
	for i := 0; i < len(cycle); i++ {
		result += cycle[i]
	}
	return result
}
