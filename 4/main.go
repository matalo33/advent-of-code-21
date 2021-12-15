package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type card struct {
	numbers   [][]int
	completed bool
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	inputs := strings.Split(string(file), "\n\n")

	theDrawStr := strings.Split(inputs[0], ",")
	theDraw := make([]int, len(theDrawStr))
	for i, str := range theDrawStr {
		theDraw[i], _ = strconv.Atoi(str)
	}

	cards := make([]card, len(inputs)-1)
	for i, cardIn := range inputs[1:] {
		rows := strings.Split(cardIn, "\n")
		thisCard := make([][]int, len(rows))
		for j, row := range rows {
			numbers := strings.Fields(row)
			cardRow := make([]int, len(numbers))
			for k, number := range numbers {
				cardRow[k], _ = strconv.Atoi(number)
			}
			thisCard[j] = cardRow
		}
		cards[i] = card{thisCard, false}
	}

	cardsCopy := make([]card, len(cards))
	copy(cardsCopy, cards)

	fmt.Printf("First winning card score: %v\n", findCardScore(theDraw, cards, true))
	fmt.Printf("Last winning card score: %v\n", findCardScore(theDraw, cardsCopy, false))
}

func findCardScore(theDraw []int, cards []card, findWinning bool) int {
	completedCards := 0
	for _, draw := range theDraw {
		for c, card := range cards {
			if !card.completed {
				for i, row := range card.numbers {
					for j, num := range row {
						if num == draw {
							card.numbers[i][j] = -1
							cardScore := isWinningCard(card, draw)
							if cardScore != 0 {
								completedCards++
								cards[c].completed = true
								if findWinning || (!findWinning && completedCards == len(cards)) {
									return cardScore
								}
							}
						}
					}
				}
			}
		}
	}
	return 0
}

func isWinningCard(card card, draw int) int {
	score := 0
	isWinner := false

	//rows
	for _, row := range card.numbers {
		tally := 0
		for _, num := range row {
			tally += num
			if num != -1 {
				score += num
			}
		}
		if tally == -len(row) {
			isWinner = true
		}
	}

	//columns
	for i := 0; i < len(card.numbers[0]); i++ {
		tally := 0
		for j := 0; j < len(card.numbers); j++ {
			tally += card.numbers[j][i]
		}
		if tally == -len(card.numbers) {
			isWinner = true
		}
	}

	if isWinner {
		return score * draw
	} else {
		return 0
	}
}
