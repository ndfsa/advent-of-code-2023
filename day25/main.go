package day24

import (
	"math"
	"math/rand"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

type Pair struct {
	so int
	ds int
}

type Graph struct {
	Vertices map[int]int
	Edges    map[Pair]int
}

func (g *Graph) Copy() Graph {
	verticesCopy := map[int]int{}
	for vertex, count := range g.Vertices {
		verticesCopy[vertex] = count
	}

	edgesCopy := map[Pair]int{}
	for vertex, value := range g.Edges {
		edgesCopy[vertex] = value
	}
	return Graph{verticesCopy, edgesCopy}

}

func fastMinCut(g Graph) (int, int) {
	if len(g.Vertices) <= 6 {
		contract(&g, 2)

		sum := 0
		for _, count := range g.Edges {
			sum += count
		}
		prod := 1
		for _, count := range g.Vertices {
			prod *= count
		}
		return sum / 2, prod
	}

	t := math.Ceil(1 + float64(len(g.Vertices))/math.Sqrt2)

	G1 := g.Copy()
	G2 := g.Copy()

	contract(&G1, int(t))
	contract(&G2, int(t))

	r1, p1 := fastMinCut(G1)
	r2, p2 := fastMinCut(G2)

	if r1 < r2 {
		return r1, p1
	}
	return r2, p2
}

// Karger's algorithm
func contract(g *Graph, t int) {
	for len(g.Vertices) > t {
		var edge Pair
		k := rand.Intn(len(g.Edges))
		for neighbor := range g.Edges {
			if k == 0 {
				edge = neighbor
				break
			}
			k--
		}

		u, v := edge.so, edge.ds
		delete(g.Edges, Pair{u, v})
		delete(g.Edges, Pair{v, u})
		newEdges := map[Pair]int{}
		for edge, count := range g.Edges {
			if edge.so == v {
				delete(g.Edges, edge)
				newEdges[Pair{u, edge.ds}] = count
			}
			if edge.ds == v {
				delete(g.Edges, edge)
				newEdges[Pair{edge.so, u}] = count
			}
		}
		for edge, count := range newEdges {
			g.Edges[edge] += count
		}
		g.Vertices[u] += g.Vertices[v]
		delete(g.Vertices, v)
	}
}

func parseInput(lines []string) (map[int]int, map[Pair]int) {
	edges := map[Pair]int{}
	nodes := map[int]int{}

	t := map[string]int{}
	count := 0

	for _, line := range lines {
		node, neighbors, _ := strings.Cut(line, ": ")
		if _, ok := t[node]; !ok {
			t[node] = count
			count++
		}
		nodes[t[node]] = 1
		for _, neighbor := range strings.Split(neighbors, " ") {
			if _, ok := t[neighbor]; !ok {
				t[neighbor] = count
				count++
			}
			edges[Pair{t[node], t[neighbor]}] = 1
			edges[Pair{t[neighbor], t[node]}] = 1
			nodes[t[neighbor]] = 1
		}
	}

	return nodes, edges
}

func Solve(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	vertices, edges := parseInput(lines)
	graph := Graph{vertices, edges}

	for {
		graphCopy := graph.Copy()
		if minCut, prod := fastMinCut(graphCopy); minCut != 3 {
			continue
		} else {
			return prod, nil
		}
	}
}
