package ds

import (
	"math"
)

// EdgesContainer defines the contract of Edges containers implementations
type EdgesContainer interface {
	IsAdjacent(i int, j int) bool
	weight(i int, j int) int
	eachAdjacent(i int, fn func(int, int))
}

// ListContainer edges linked list implementation
type ListContainer struct {
	edges []LinkedList
}

// NewListContainer creates a new linked list to store edges
func NewListContainer(verticesCount int) ListContainer {
	edges := make([]LinkedList, verticesCount)

	for i := 0; i < verticesCount; i++ {
		edges[i] = NewLinkedList()
	}

	return ListContainer{
		edges: edges,
	}
}

// InsertEdge adds a new edge to the container
func (l *ListContainer) InsertEdge(i int, j int, weight int) {
	l.edges[i].Insert(j, weight)
}

func (l ListContainer) weight(i int, j int) int {
	edge := l.edges[i].Search(j)
	if edge == nil {
		return -1
	}
	weight := edge.Data()
	return weight.(int)
}

// IsAdjacent verify if two vertices are adjacent
func (l ListContainer) IsAdjacent(i int, j int) bool {
	edge := l.edges[i].Search(j)
	return edge != nil
}

// eachAdjacent executes a given function for each adjacent of the vertex i
func (l ListContainer) eachAdjacent(i int, fn func(int, int)) {
	l.edges[i].Each(func(j int, node LinkedListNode) {
		fn(node.Key(), l.weight(i, node.Key()))
	})
}

// Vertex vertex implementation
type Vertex struct {
	key  int
	data interface{}
}

// NewVertex creates a new vertex
func NewVertex(key int, data interface{}) Vertex {
	return Vertex{
		key:  key,
		data: data,
	}
}

// Graph graph implementation
type Graph struct {
	vertices []Vertex
	edges    EdgesContainer
}

// NewGraph creates a new graph given vertices and edges
func NewGraph(vertices []Vertex, edges EdgesContainer) Graph {
	return Graph{
		vertices: vertices,
		edges:    edges,
	}
}

// Weight returns the Weight for the i,j edge
func (g *Graph) Weight(i int, j int) int {
	return g.edges.weight(i, j)
}

func (g *Graph) eachAdjacent(i int, fn func(int, int)) {
	g.edges.eachAdjacent(i, fn)
}

// MinimumPath returns the minimum path and distance from root to every vertex (uses Dijkstra)
func (g *Graph) MinimumPath(root int) ([]int, []int) {
	dist := make([]int, len(g.vertices))
	pred := make([]int, len(g.vertices))

	for v := 0; v < len(g.vertices); v++ {
		dist[v] = math.MaxInt32
		pred[v] = -1
	}

	dist[root] = 0

	sptSet := make([]bool, len(g.vertices))

	for u := root; u >= 0 && dist[u] != math.MaxInt32; u = minDistance(dist, sptSet) {
		sptSet[u] = true

		g.eachAdjacent(u, func(v int, weight int) {
			if dist[v] > dist[u]+weight {
				dist[v] = dist[u] + weight
				pred[v] = u
			}
		})
	}

	return pred, dist
}

// minDistance returns the index of the minimum not verified distance
func minDistance(dist []int, sptSet []bool) int {
	min := -1
	for i := 0; i < len(dist); i++ {
		if !sptSet[i] {
			if (min == -1) || (dist[i] < dist[min]) {
				min = i
			}
		}
	}
	return min
}

// MinimumPathTo returns the minimum path and distance between two vertices (uses Dijkstra)
func (g *Graph) MinimumPathTo(root int, target int) ([]int, int) {
	pred, dist := g.MinimumPath(root)
	path := NewStack(len(g.vertices))

	parent := pred[target]
	current := target
	count := 0
	for parent != -1 {
		path.Push(current)
		current = parent
		parent = pred[current]
		count++
	}

	path.Push(root)

	v := make([]int, path.Count())
	i := 0
	for !path.IsEmpty() {
		v[i], _ = path.Pop()
		i++
	}

	return v, dist[target]
}
