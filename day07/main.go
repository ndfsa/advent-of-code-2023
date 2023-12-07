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

func (h *Hand) calculateType() {
	mp := make(map[int]int, 0)
	for _, card := range h.Cards {
		if _, ok := mp[card]; ok {
			mp[card]++
		} else {
			mp[card] = 1
		}
	}
	vals := make([]int, 0, len(mp))
	for _, v := range mp {
		vals = append(vals, v)
	}
	switch len(mp) {
	case 1:
		h.Type = FIVE_OF_A_KIND
	case 2:
		if slices.Max(vals) == 4 {
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

func parseInput(lines []string) []Hand {
	res := make([]Hand, 0, len(lines))
	for _, line := range lines {
		var cards string
		var bid int

		fmt.Sscanf(line, "%s %d", &cards, &bid)

		converted := make([]int, 0)
		for _, card := range []byte(cards) {
			converted = append(converted, cardValue[card])
		}
		res = append(res, Hand{Bid: bid, Cards: converted})
	}
	for i := range res {
		res[i].calculateType()
	}

	return res
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	hands := parseInput(lines)
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Type == hands[j].Type &&
			hands[i].lt(hands[j]) ||
			hands[i].Type < hands[j].Type
	})
	res := 0
	for i, hand := range hands {
		res += (i + 1) * hand.Bid
	}

	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	_, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	return 0, nil
}
