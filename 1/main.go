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

	var measurements []int

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		measurements = append(measurements, i)
	}

	fmt.Printf("Measurements larger: %v\n", countLarger(measurements))
	fmt.Printf("Measurements larger with sliding window: %v\n", countSlidingWindowsLarger(measurements))
}

func countLarger(measurements []int) (count int) {
	current := measurements[0]

	for _, i := range measurements[1:] {
		if i > current {
			count++
		}
		current = i
	}

	return count
}

func countSlidingWindowsLarger(measurements []int) (count int) {
	current := measurements[0] + measurements[1] + measurements[2]

	for i := 1; i < len(measurements)-2; i++ {
		this := measurements[i] + measurements[i+1] + measurements[i+2]
		if this > current {
			count++
		}
		current = this
	}

	return count
}
