package day06

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

func parseInput1(input string) ([]int, []int) {
	re := regexp.MustCompile(`(\d+)`)
	time := make([]int, 0)
	distance := make([]int, 0)
	lines := strings.Split(input, "\n")
	for _, match := range re.FindAllString(lines[0], -1) {
		num, _ := strconv.Atoi(match)
		time = append(time, num)
	}
	for _, match := range re.FindAllString(lines[1], -1) {
		num, _ := strconv.Atoi(match)
		distance = append(distance, num)
	}
	return time, distance
}

func solution(t, d int) (int, int) {
	tFloat := float64(t)
	dFloat := float64(d)

	x1 := (tFloat + math.Sqrt(math.Pow(tFloat, 2.0)-4.0*dFloat)) / 2.0
	x2 := (tFloat - math.Sqrt(math.Pow(tFloat, 2.0)-4.0*dFloat)) / 2.0

	return int(math.Ceil(x1)), int(math.Floor(x2))
}

func SolvePart1(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	time, distance := parseInput1(input)

	res := 1
	for i := 0; i < len(time); i++ {
		x1, x2 := solution(time[i], distance[i])
		res *= x1 - x2 - 1
	}

	return res, nil
}

func parseInput2(input string) (int, int) {
	re := regexp.MustCompile(`(\d+)`)
	lines := strings.Split(input, "\n")

	timeStr := ""
	distanceStr := ""
	for _, match := range re.FindAllString(lines[0], -1) {
		timeStr += match
	}
	for _, match := range re.FindAllString(lines[1], -1) {
		distanceStr += match
	}

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)

	return time, distance
}

func SolvePart2(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	time, distance := parseInput2(input)
	x1, x2 := solution(time, distance)

	return x1 - x2 - 1, nil
}
