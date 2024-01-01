package day14

import (
	"strconv"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Instruction struct {
	box   int
	label string
	fl    int
}

func parseInput(input string) []Instruction {
	instructions := []Instruction{}
	for _, inst := range strings.Split(input, ",") {
		if label, flString, ok := strings.Cut(inst, "="); ok {
			fl, _ := strconv.Atoi(flString)
			box := hash(label)

			instructions = append(instructions, Instruction{box, label, fl})
		} else {
			label, _, _ := strings.Cut(inst, "-")
			box := hash(label)

			instructions = append(instructions, Instruction{box, label, 0})
		}
	}
	return instructions
}

// Holiday ASCII String Helper
func hash(input string) int {
	res := 0
	for _, ch := range input {
		res += int(ch)
		res *= 17
		res %= 256
	}
	return res
}

func SolvePart1(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	words := []string{}
	for _, word := range strings.Split(input, ",") {
		words = append(words, word)
	}

	res := 0
	for _, word := range words {
		res += hash(word)
	}

	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	instructions := parseInput(lines)
	boxes := make([][]*Instruction, 256)

	for idx := range instructions {
		boxIdx := instructions[idx].box
		if boxes[boxIdx] == nil {
			boxes[boxIdx] = make([]*Instruction, 0)
		}

		found := -1
		i := 0
		for i < len(boxes[boxIdx]) {
			if boxes[boxIdx][i].label == instructions[idx].label {
				found = i
				break
			}
			i++
		}

		if instructions[idx].fl == 0 && found != -1 {
			boxes[boxIdx] = append(boxes[boxIdx][:found], boxes[boxIdx][found+1:]...)
		} else if found != -1 {
			boxes[boxIdx][found] = &instructions[idx]
		} else if instructions[idx].fl != 0{
			boxes[boxIdx] = append(boxes[boxIdx], &instructions[idx])
		}
	}

	res := 0
	for _, box := range boxes {
		if len(box) == 0 {
			continue
		}
		for i, instruction := range box {
			fp := (instruction.box + 1) * (i + 1) * (instruction.fl)
			res += fp
		}
	}

	return res, nil
}
