package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
	"golang.org/x/exp/slices"
)

type Range struct {
	start int
	end   int
}

type RangeMap struct {
	rang   Range
	offset int
}

func (r Range) contains(val int) bool {
	return r.start <= val && val <= r.end
}

func (r RangeMap) execute(val int) int {
	return val + r.offset
}

type AlmanacMap struct {
	ranges []RangeMap
}

func parseInput(input string) ([]int, []AlmanacMap) {
	chunks := strings.Split(input, "\n\n")

	seeds := make([]int, 0)
	for _, seed := range strings.Split(chunks[0], " ")[1:] {
		num, _ := strconv.Atoi(seed)
		seeds = append(seeds, num)
	}

	almanac := make([]AlmanacMap, 0)
	for _, chunk := range chunks[1:] {
		almanacMap := AlmanacMap{}
		for _, rangeStr := range strings.Split(chunk, "\n")[1:] {
			var dest, source, length int
			fmt.Sscanf(rangeStr, "%d %d %d",
				&dest,
				&source,
				&length)
			almanacMap.ranges = append(almanacMap.ranges,
				RangeMap{rang: Range{start: source, end: source + length - 1}, offset: dest - source})
		}

		almanac = append(almanac, almanacMap)
	}

	return seeds, almanac
}

func createFunction(almanac []AlmanacMap) func(int) int {
	return func(i int) int {
		for _, almanacMap := range almanac {
			for _, rangeMap := range almanacMap.ranges {
				if rangeMap.rang.contains(i) {
					i = rangeMap.execute(i)
					break
				}
			}
		}
		return i
	}
}

func SolvePart1(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}
	seeds, almanac := parseInput(input)
	f := createFunction(almanac)

	for idx := range seeds {
		seeds[idx] = f(seeds[idx])
	}

	return slices.Min(seeds), nil
}

func SolvePart2(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	seeds, almanac := parseInput(input)
	for i, j := 0, len(almanac)-1; i < j; i, j = i+1, j-1 {
		almanac[i], almanac[j] = almanac[j], almanac[i]
	}

	for i := range almanac {
		for j := range almanac[i].ranges {
			offset := almanac[i].ranges[j].offset
			almanac[i].ranges[j].rang.start += offset
			almanac[i].ranges[j].rang.end += offset
			almanac[i].ranges[j].offset = -offset
		}
	}

	seedsRanges := make([]Range, 0)
	for i := range seeds {
		if i%2 == 1 {
			seedsRanges = append(seedsRanges, Range{start: seeds[i-1], end: seeds[i-1] + seeds[i] - 1})
		}
	}

	fInv := createFunction(almanac)
	for i := 0; i < 1<<31; i++ {
		seed := fInv(i)
		for _, r := range seedsRanges {
			if r.contains(seed) {
				return i, nil
			}
		}
	}

	return 0, nil
}
