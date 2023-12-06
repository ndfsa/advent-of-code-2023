package day04

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Card struct {
	winning []int
	played  []int
	points  int
}

func parseInput(lines []string) []Card {
	re := regexp.MustCompile(`(\d+)`)

	cards := make([]Card, 0, len(lines))
	points := 0
	for _, line := range lines {
		lists := strings.Split(strings.Split(line, ":")[1], "|")

		winning := make([]int, 0)
		for _, w := range re.FindAllStringSubmatch(lists[0], -1) {
			num, _ := strconv.Atoi(w[0])
			winning = append(winning, num)
		}

		played := make([]int, 0)
		for _, w := range re.FindAllStringSubmatch(lists[1], -1) {
			num, _ := strconv.Atoi(w[0])
			played = append(played, num)
		}

		cards = append(cards, Card{winning, played, points})
	}

	for idx, card := range cards {
	winning:
		for _, w := range card.winning {
			for _, p := range card.played {
				if w == p {
					cards[idx].points++
					continue winning
				}
			}
		}
	}
	return cards
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	res := 0
	for _, card := range parseInput(lines) {
		if card.points != 0 {
			res += 1 << (card.points - 1)
		}
	}

	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	input, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	cards := parseInput(input)
	copies := make([]int, len(cards))
	for idx, card := range cards {
		for offset := 1; offset <= card.points; offset++ {
			copies[idx+offset] += copies[idx] + 1
		}
	}

	res := len(cards)
	for _, c := range copies {
		res += c
	}

	return res, nil
}
