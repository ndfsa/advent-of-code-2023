package day07

import (
	"fmt"
	"sort"

	"github.com/ndfsa/advent-of-code-2023/util"
	"golang.org/x/exp/slices"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

var cardValue = map[byte]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

var cardValueModified = map[byte]int{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type Hand struct {
	Bid   int
	Cards []int
	Type  int
}

func (h Hand) lt(other Hand) bool {
	for i := 0; i < 5; i++ {
		if h.Cards[i] != other.Cards[i] {
			return h.Cards[i] < other.Cards[i]
		}
	}
	return false
}

func (h *Hand) calculateType(joker bool) {
	mp := make(map[int]int, 0)
	for _, card := range h.Cards {
		if _, ok := mp[card]; ok {
			mp[card]++
		} else {
			mp[card] = 1
		}
	}
	if joker {
		if jk, ok := mp[cardValueModified['J']]; ok {
			delete(mp, cardValueModified['J'])
			maxValKey, maxVal := -1, 0
			for k, v := range mp {
				if v > maxVal {
					maxValKey, maxVal = k, v
				}
			}
			mp[maxValKey] += jk
		}
	}
	keys := make([]int, 0, len(mp))
	vals := make([]int, 0, len(mp))
	for k, v := range mp {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	maxVal := slices.Max(vals)
	switch len(mp) {
	case 1:
		h.Type = FIVE_OF_A_KIND
	case 2:
		if maxVal == 4 {
			h.Type = FOUR_OF_A_KIND
		} else {
			h.Type = FULL_HOUSE
		}
	case 3:
		if slices.Max(vals) == 3 {
			h.Type = THREE_OF_A_KIND
		} else {
			h.Type = TWO_PAIR
		}
	case 4:
		h.Type = ONE_PAIR
	case 5:
		h.Type = HIGH_CARD
	}
}

func parseInput(lines []string, joker bool) []Hand {
	res := make([]Hand, 0, len(lines))
	for _, line := range lines {
		var cards string
		var bid int

		fmt.Sscanf(line, "%s %d", &cards, &bid)

		converted := make([]int, 0)
		for _, card := range []byte(cards) {
			if joker {
				converted = append(converted, cardValueModified[card])
			} else {
				converted = append(converted, cardValue[card])
			}
		}
		res = append(res, Hand{Bid: bid, Cards: converted})
	}
	return res
}

func solveProblem(lines []string, joker bool) int {
	hands := parseInput(lines, false)
	for i := range hands {
		hands[i].calculateType(false)
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Type == hands[j].Type &&
			hands[i].lt(hands[j]) ||
			hands[i].Type < hands[j].Type
	})
	res := 0
	for i, hand := range hands {
		res += (i + 1) * hand.Bid
	}
	return res
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	return solveProblem(lines, false), nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	return solveProblem(lines, true), nil
}
