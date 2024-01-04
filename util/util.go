package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func ReadFile(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	res := string(file)
	return strings.Trim(res, "\n"), nil
}

func ReadFileSplit(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func RunSolution[T comparable](
	t *testing.T,
	solution func(string) (T, error),
	filePath string,
	expected T) {

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		t.Skip(err)
	}

	res, err := solution(filePath)

	if err != nil {
		t.Fatal(err)
	}

	if res != expected {
		t.Fatal(fmt.Sprintf("\nincorrect: %v\nexpected: %v", res, expected))
	}
	t.Logf("res: %v", res)
}

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(integers []int) int {
	if len(integers) == 2 {
		a, b := integers[0], integers[1]
		return a * b / GCD(a, b)
	}

	return LCM([]int{integers[0], LCM(integers[1:])})
}

func AbsDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

type Vec2 struct {
	Row int
	Col int
}

var DIR_UP = Vec2{Row: -1, Col: 0}
var DIR_DOWN = Vec2{Row: 1, Col: 0}
var DIR_RIGHT = Vec2{Row: 0, Col: 1}
var DIR_LEFT = Vec2{Row: 0, Col: -1}

func (p Vec2) Add(other Vec2) Vec2 {
	return Vec2{Row: p.Row + other.Row, Col: p.Col + other.Col}
}

func (p Vec2) AddMult(other Vec2, mag int) Vec2 {
	return Vec2{Row: p.Row + other.Row * mag, Col: p.Col + other.Col * mag}
}

func (p Vec2) HammiltonDist(other Vec2) int {
	sus := p.Sus(other)
	res := 0

	if sus.Row < 0 {
		res += -sus.Row
	} else {
		res += sus.Row
	}

	if sus.Col < 0 {
		res += -sus.Col
	} else {
		res += sus.Col
	}

	return res
}

func (p Vec2) Sus(other Vec2) Vec2 {
	return Vec2{Row: p.Row - other.Row, Col: p.Col - other.Col}
}

func SlicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
