package day23

import (
	"container/heap"
	"math"
	"slices"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Graph struct {
	Vertices []util.Vec2
	Edges    map[util.Vec2]map[util.Vec2]int
}

var (
	DIR_UP    = util.DIR_V2_NEG_X
	DIR_DOWN  = util.DIR_V2_POS_X
	DIR_RIGHT = util.DIR_V2_POS_Y
	DIR_LEFT  = util.DIR_V2_NEG_Y
)

func parseInput(lines []string) (Graph, [][]byte) {
	field := [][]byte{}
	for _, line := range lines {
		field = append(field, []byte(line))
	}

	vertices := []util.Vec2{}
	vertices = append(vertices, util.Vec2{X: 0, Y: strings.Index(lines[0], ".")})
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if field[i][j] == '#' {
				continue
			}

			pos := util.Vec2{X: i, Y: j}
			if len(pos.Neighbors(func(v, dir util.Vec2) bool {
				if v.X < 0 ||
					v.X >= len(field) ||
					v.Y < 0 ||
					v.Y >= len(field[0]) {
					return false
				}

				return field[v.X][v.Y] != '#'
			})) > 2 {
				vertices = append(vertices, pos)
			}
		}
	}
	vertices = append(vertices, util.Vec2{
		X: len(lines) - 1,
		Y: strings.Index(lines[len(lines)-1], ".")})

	edges := map[util.Vec2]map[util.Vec2]int{}
	for _, vert := range vertices {
		edges[vert] = map[util.Vec2]int{}
	}
	return Graph{Vertices: vertices, Edges: edges}, field
}

func (g *Graph) FillEdges(
	field [][]byte,
	accept func(v util.Vec2, dir util.Vec2) bool) {

	for _, vert := range g.Vertices {
		queue := []util.State[util.Vec2]{{Vertex: vert, Cost: 0}}
		visited := map[util.Vec2]struct{}{}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if slices.Contains(g.Vertices, current.Vertex) && current.Vertex != vert {
				g.Edges[vert][current.Vertex] = current.Cost
				continue
			}

			neighbors := current.Vertex.Neighbors(func(v, dir util.Vec2) bool {
				if _, ok := visited[v]; ok {
					return false
				}
				return accept(v, dir)
			})

			for _, neighbor := range neighbors {
				queue = append(queue, util.State[util.Vec2]{
					Vertex: neighbor,
					Cost:   current.Cost + 1})
				visited[neighbor] = struct{}{}
			}
		}
	}
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	graph, field := parseInput(lines)
	graph.FillEdges(field, func(v, dir util.Vec2) bool {
		if v.X < 0 ||
			v.X >= len(field) ||
			v.Y < 0 ||
			v.Y >= len(field[0]) {
			return false
		}

		ch := field[v.X][v.Y]
		if ch == '#' {
			return false
		}

		switch dir {
		case DIR_UP:
			return ch != 'v'
		case DIR_DOWN:
			return ch != '^'
		case DIR_RIGHT:
			return ch != '<'
		case DIR_LEFT:
			return ch != '>'
		default:
			return false
		}
	})

	start := graph.Vertices[0]
	end := graph.Vertices[len(graph.Vertices)-1]

	queue := &util.StateHeap[util.Vec2]{{Vertex: start, Cost: 0}}
	heap.Init(queue)

	res := 0
	for queue.Len() > 0 {
		current := heap.Pop(queue).(util.State[util.Vec2])

		if current.Vertex == end && current.Cost < res {
			res = current.Cost
		}

		for neighbor, cost := range graph.Edges[current.Vertex] {
			heap.Push(queue, util.State[util.Vec2]{Vertex: neighbor, Cost: current.Cost - cost})
		}
	}

	return -res, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	graph, field := parseInput(lines)
	graph.FillEdges(field, func(v, dir util.Vec2) bool {
		if v.X < 0 ||
			v.X >= len(field) ||
			v.Y < 0 ||
			v.Y >= len(field[0]) {
			return false
		}

		return field[v.X][v.Y] != '#'
	})

	start := graph.Vertices[0]
	end := graph.Vertices[len(graph.Vertices)-1]

	visited := []util.Vec2{}

	var dfs func(util.Vec2) int
	dfs = func(current util.Vec2) int {
		if current == end {
			return 0
		}

		maxCost := math.MinInt

		visited = append(visited, current)

		for neighbor, cost := range graph.Edges[current] {
			if !slices.Contains(visited, neighbor) {
				maxCost = max(maxCost, dfs(neighbor)+cost)
			}
		}

		visited = visited[:len(visited)-1]

		return maxCost
	}

	return dfs(start), nil
}
