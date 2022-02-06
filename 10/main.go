package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var navs []string
	for scanner.Scan() {
		navs = append(navs, scanner.Text())
	}

	errorScore, completionScore := getScores(navs)

	fmt.Printf("Syntax error score: %v\n", errorScore)
	fmt.Printf("Autocomplete score: %v\n", completionScore)
}

func getScores(navs []string) (int, int) {
	var errorScore int
	var completionScores []int
nextWord:
	for _, nav := range navs {
		var closeList []rune
		for _, this := range nav {
			closing, err := findClosingSymbol(this)
			if err == nil {
				closeList = append(closeList, closing)
			} else {
				if this == closeList[len(closeList)-1] {
					closeList = closeList[:len(closeList)-1]
				} else {
					errorScore += scores[this]
					continue nextWord
				}
			}
		}

		// Getting here means the nav was incomplete
		reverseCloseList := reverseSlice(closeList)
		completionScore := 0
		for _, i := range reverseCloseList {
			completionScore = completionScore * 5
			completionScore += completionScoreValue(i)
		}
		completionScores = append(completionScores, completionScore)
	}

	sort.Ints(completionScores)
	return errorScore, completionScores[(len(completionScores)-1)/2]
}

func findClosingSymbol(symbol rune) (rune, error) {
	switch symbol {
	case '(':
		return ')', nil
	case '[':
		return ']', nil
	case '{':
		return '}', nil
	case '<':
		return '>', nil
	default:
		return -1, errors.New("err")
	}
}

func completionScoreValue(symbol rune) int {
	switch symbol {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	default:
		return 0
	}
}

func reverseSlice(a []rune) []rune {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}
