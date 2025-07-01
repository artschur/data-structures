package graphs

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	Vertices map[int]*Vertex
}

type Vertex struct {
	Val   int
	Edges map[int]*Edge
}
type Edge struct {
	Weight int
	To     *Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[int]*Vertex),
	}
}

func (g *Graph) AddVertex(val int) {
	if _, exists := g.Vertices[val]; !exists {
		g.Vertices[val] = &Vertex{
			Val:   val,
			Edges: make(map[int]*Edge),
		}
	}
}

func (g *Graph) CreateConnection(from, to int) {
	g.AddVertex(from)
	g.AddVertex(to)
	g.Vertices[from].Edges[to] = &Edge{
		Weight: 1,
		To:     g.Vertices[to],
	}
}

func (g *Graph) GetNeighbours(vertexKey int) []*Vertex {
	if vertex, exists := g.Vertices[vertexKey]; exists {
		neighbours := make([]*Vertex, 0, len(vertex.Edges))
		for _, edge := range vertex.Edges {
			neighbours = append(neighbours, edge.To)
		}
		return neighbours
	}
	return nil
}

func MountGraph() *Graph {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	graph := NewGraph()
	for reader.Scan() {
		line := reader.Text()
		if line == "-1,-1" {
			break
		}
		firstNode, lastNode := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		firstNodeInt, _ := strconv.Atoi(firstNode)
		lastNodeInt, _ := strconv.Atoi(lastNode)

		graph.CreateConnection(firstNodeInt, lastNodeInt)
	}

	return graph
}

func SolutionDependencyGraph() {
	graph := MountGraph()

	neighbours := graph.GetNeighbours(1)
	for _, neighbour := range neighbours {
		println(neighbour.Val) // Should print 2
	}
}
