package day08

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Node struct {
	left  *Node
	right *Node
	name  string
}

func (n Node) isFinal() bool {
	return strings.HasSuffix(n.name, "Z")
}

func (n Node) String() string {
	return fmt.Sprintf("%s", n.name)
}

type Graph struct {
	nodes []Node
	names map[string]int
}

func (g Graph) getStartNodes() []*Node {
	res := make([]*Node, 0)

	for k, v := range g.names {
		if strings.HasSuffix(k, "A") {
			res = append(res, &g.nodes[v])
		}
	}

	return res
}

func (g *Graph) getNodeByName(name string) *Node {
	idx := g.names[name]
	return &g.nodes[idx]
}

func parseInput(input []string) (string, Graph) {
	nodes := make([]Node, 0, len(input)-2)
	names := make(map[string]int, len(input)-2)

	for idx, line := range input[2:] {
		name := strings.Split(line, " = ")[0]

		nodes = append(nodes, Node{name: name})
		names[name] = idx
	}

	graph := Graph{nodes, names}

	re := regexp.MustCompile(`([A-Z0-9]{3})`)
	for _, line := range input[2:] {
		matches := re.FindAllString(line, -1)
		curr := graph.getNodeByName(matches[0])
		curr.left = graph.getNodeByName(matches[1])
		curr.right = graph.getNodeByName(matches[2])
	}

	return input[0], graph
}

func SolvePart1(filePath string) (int, error) {
	input, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	program, graph := parseInput(input)
	progLength := len(program)

	count, pc := 0, 0
	end := graph.getNodeByName("ZZZ")
	curr := graph.getNodeByName("AAA")

	for curr != end {
		switch program[pc] {
		case 'R':
			curr = curr.right
		case 'L':
			curr = curr.left
		}

		count++
		pc = (pc + 1) % progLength
	}

	return count, nil
}

func SolvePart2(filePath string) (int, error) {
	input, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	program, graph := parseInput(input)
	progLength := len(program)

	startNodes := graph.getStartNodes()

	dist := make([]int, 0)
	for _, node := range startNodes {
		pc, count := 0, 0
		for !node.isFinal() {
			switch program[pc] {
			case 'R':
				node = node.right
			case 'L':
				node = node.left
			}

			count++
			pc = (pc + 1) % progLength
		}
		dist = append(dist, count)
	}

	return util.LCM(dist), nil
}
