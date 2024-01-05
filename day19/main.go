package day19

import (
	"fmt"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type PartRange struct {
	ranges map[byte][2]int
}

func (p *PartRange) adjust(r Rule, value bool) {
	if r.operation == 0 {
		return
	}

	rn := p.ranges[r.parameter]
	if value {
		switch r.operation {
		case '>':
			if rn[0] <= r.value {
				rn[0] = r.value + 1
			}
		case '<':
			if rn[1] >= r.value {
				rn[1] = r.value - 1
			}
		}
	} else {
		switch r.operation {
		case '>':
			if rn[1] > r.value {
				rn[1] = r.value
			}
		case '<':
			if rn[0] < r.value {
				rn[0] = r.value
			}
		}
	}
	p.ranges[r.parameter] = rn
}

type Part struct {
	parameters map[byte]int
}

func (p Part) score() int {
	res := 0
	for _, val := range p.parameters {
		res += val
	}
	return res
}

type Rule struct {
	parameter byte
	operation byte
	value     int
	next      string
}

type Workflow struct {
	rules []Rule
}

type RulePos struct {
	name string
	pos  int
}

func (w Workflow) execute(p Part) string {
	for _, rule := range w.rules {
		if rule.operation == 0 {
			return rule.next
		}

		var partValue int = p.parameters[rule.parameter]

		switch rule.operation {
		case '>':
			if partValue > rule.value {
				return rule.next
			}
		case '<':
			if partValue < rule.value {
				return rule.next
			}
		}
	}
	return ""
}

func parseInput(input string) (map[string]Workflow, []Part) {
	workflowsInput, partsInput, _ := strings.Cut(input, "\n\n")

	workflows := map[string]Workflow{}
	for _, workflowStr := range strings.Split(workflowsInput, "\n") {
		name, rulesStr, _ := strings.Cut(workflowStr, "{")

		rules := []Rule{}
		rulesSplit := strings.Split(rulesStr[:len(rulesStr)-1], ",")
		for _, ruleStr := range rulesSplit[:len(rulesSplit)-1] {
			var newRule Rule
			fmt.Sscanf(ruleStr, "%c%c%d:%s",
				&newRule.parameter,
				&newRule.operation,
				&newRule.value,
				&newRule.next)

			rules = append(rules, newRule)
		}
		rules = append(rules, Rule{next: rulesSplit[len(rulesSplit)-1]})

		workflows[name] = Workflow{rules}
	}

	parts := []Part{}
	for _, partStr := range strings.Split(partsInput, "\n") {
		x, m, a, s := 0, 0, 0, 0
		fmt.Sscanf(partStr, "{x=%d,m=%d,a=%d,s=%d}",
			&x,
			&m,
			&a,
			&s)
		parts = append(parts, Part{parameters: map[byte]int{
			'x': x,
			'm': m,
			'a': a,
			's': s}})
	}

	return workflows, parts
}

func SolvePart1(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	workflows, parts := parseInput(input)

	curr := "in"
	res := 0
	for _, part := range parts {
	workflow:
		for {
			if curr == "A" {
				res += part.score()
				break
			} else if curr == "R" {
				break workflow
			} else {
				curr = workflows[curr].execute(part)
			}
		}
		curr = "in"
	}

	return res, nil
}

func SolvePart2(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	workflows, _ := parseInput(input)

	rulePositions := []RulePos{}
	for name, workflow := range workflows {
		for i, rule := range workflow.rules {
			if rule.next == "A" {
				rulePositions = append(rulePositions, RulePos{name, i})
			}
		}
	}

	acceptRanges := []PartRange{}
	for _, rulePos := range rulePositions {
		partRange := PartRange{ranges: map[byte][2]int{
			'x': {1, 4000},
			'm': {1, 4000},
			'a': {1, 4000},
			's': {1, 4000}}}

		for {
			workflow := workflows[rulePos.name]
			rule := workflow.rules[rulePos.pos]

			partRange.adjust(rule, true)
			for i := rulePos.pos - 1; i >= 0; i-- {
				partRange.adjust(workflow.rules[i], false)
			}

			if rulePos.name == "in" {
				acceptRanges = append(acceptRanges, partRange)
				break
			}

		out:
			for name, workflow := range workflows {
				for i, rule := range workflow.rules {
					if rule.next == rulePos.name {
						rulePos.name = name
						rulePos.pos = i
						break out
					}
				}
			}
		}
	}

	res := 0
	for _, partRange := range acceptRanges {
		partial := 1
		for _, rn := range partRange.ranges {
			partial *= rn[1] - rn[0] + 1
		}
		res += partial
	}

	return res, nil
}
