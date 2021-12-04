package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	amount    int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var instructions []instruction

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		amount, _ := strconv.Atoi(words[1])
		instructions = append(instructions, instruction{
			words[0],
			amount})
	}

	fmt.Printf("Final position value: %v\n", calcFinalPosition(instructions))
	fmt.Printf("Final position value with aiming: %v\n", calcFinalPositionWithAim(instructions))
}

func calcFinalPosition(instructions []instruction) int {
	x, y := 0, 0
	for _, instruction := range instructions {
		if instruction.direction == "forward" {
			x += instruction.amount
		} else if instruction.direction == "down" {
			y += instruction.amount
		} else if instruction.direction == "up" {
			y -= instruction.amount
		}
	}
	return x * y
}

func calcFinalPositionWithAim(instructions []instruction) int {
	aim, x, y := 0, 0, 0
	for _, instruction := range instructions {
		if instruction.direction == "forward" {
			x += instruction.amount
			y += (instruction.amount * aim)
		} else if instruction.direction == "down" {
			aim += instruction.amount
		} else if instruction.direction == "up" {
			aim -= instruction.amount
		}
	}
	return x * y
}
