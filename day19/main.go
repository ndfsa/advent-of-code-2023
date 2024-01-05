package day19

import (
	"fmt"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Part struct {
	x int
	m int
	a int
	s int
}

func (p Part) score() int {
	return p.x + p.m + p.a + p.s
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

func (w Workflow) execute(p Part) string {
	for _, rule := range w.rules {
		if rule.operation == 0 {
			return rule.next
		}

		var partValue int
		switch rule.parameter {
		case 'x':
			partValue = p.x
		case 'm':
			partValue = p.m
		case 'a':
			partValue = p.a
		case 's':
			partValue = p.s
		}

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
		newPart := Part{}
		fmt.Sscanf(partStr, "{x=%d,m=%d,a=%d,s=%d}",
			&newPart.x,
			&newPart.m,
			&newPart.a,
			&newPart.s)
		parts = append(parts, newPart)
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

	// find all rules where next = A

	// for each accept go through the rules and narrow down an interval:

	// 1. the current rule should be true
	// 2. all the preceding rules need to be false
	
	// 3. if the workflow is "in" count the combinations and continue

	// 4. if it is not, find the rules that lead to the current workflow, on first glance it seems
	// that each rule has exactly one preceding rule

	// 5. the current workflow and rule are now the workflow and rule that lead here, go to 1.

	// add up the combinations, you may need a bigint

	parseInput(input)
	return 0, nil
}
