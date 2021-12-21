package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	line := strings.Split(string(file), "\n")
	input := strings.Split(line[0], ",")

	crabs := make([]int, len(input))
	for i, f := range input {
		crabs[i], _ = strconv.Atoi(f)
	}

	fmt.Printf("Best horizontal position: %v\n", findHorizontalPosition(crabs, true))
	fmt.Printf("Best horizontal position with costly fuel: %v\n", findHorizontalPosition(crabs, false))

}

func findHorizontalPosition(crabs []int, simpleFuel bool) int {
	min, max := getMinMax(crabs)
	bestFuel := math.MaxInt
	for target := min; target <= max; target++ {
		fuel := 0
		for _, crab := range crabs {
			res := target - crab
			if res < 0 {
				if simpleFuel {
					fuel += -res
				} else {
					fuel += fuelCost(-res)
				}
			} else {
				if simpleFuel {
					fuel += res
				} else {
					fuel += fuelCost(res)
				}
			}
		}
		if fuel < bestFuel {
			bestFuel = fuel
		}
	}
	return bestFuel
}

func fuelCost(fuel int) (cost int) {
	for i := 0; i <= fuel; i++ {
		cost += i
	}
	return cost
}

func getMinMax(arr []int) (int, int) {
	var max int = arr[0]
	var min int = arr[0]
	for _, val := range arr {
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}
	}
	return min, max
}
