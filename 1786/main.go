package main

import "fmt"

func main() {
	edges := [][]int{
		{1, 2, 3},
		{1, 3, 3},
		{2, 3, 1},
		{1, 4, 2},
		{5, 2, 2},
		{3, 5, 1},
		{5, 4, 10},
	}

	graph := TurnIntoGraph(5, edges)
	fmt.Print("graph: \n", graph)

	neighbours := graph.GetNeighbours(2)
	fmt.Print("\nneighbours of node 2: ", neighbours)
	neighbours = graph.GetNeighbours(1)
	fmt.Print("\nneighbours of node 1: ", neighbours)
	neighbours = graph.GetNeighbours(5)
	fmt.Print("\nneighbours of node 5: ", neighbours)

	degree := graph.GetDegree(4)
	fmt.Print("\ndegree of node 4: ", degree)
	degree = graph.GetDegree(2)
	fmt.Print("\ndegree of node 2: ", degree)

	fmt.Println()
	matrix := graph.TurnIntoMatrix()
	for i := 0; i < len(matrix); i++ {
		fmt.Printf("%d: ", i+1)
		fmt.Printf(" %d \n", matrix[i])
	}

	visitedNodes := make(map[int]bool, len(graph))
	graph.RecursiveDepthFirstSearch(1, visitedNodes)
	visitedNodesIt := make(map[int]bool, len(graph))
	graph.IterativeDepthFirstSearch(1, visitedNodesIt)
}

func (g Graph) RecursiveDepthFirstSearch(initialNode int, visitedNodes map[int]bool) {
	visitedNodes[initialNode] = true
	fmt.Println("recursive visited node: ", initialNode)
	neighbours := g.GetNeighbours(initialNode)
	for _, neighbour := range neighbours {
		if !visitedNodes[neighbour] {
			g.RecursiveDepthFirstSearch(neighbour, visitedNodes)
		}
	}
}

func (g Graph) IterativeDepthFirstSearch(initialNode int, visitedNodes map[int]bool) {
	stack := []int{initialNode}
	visitedNodes[initialNode] = true

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println("iterative visited node: ", node)

		for _, neighbour := range g.GetNeighbours(node) {
			if !visitedNodes[neighbour] {
                visitedNodes[neighbour] = true
				stack = append(stack, neighbour)
			}
		}
	}
}

// Graph G = (V, E) is a set of vertices V and edges E,
// where each edge (u,v) is a conection between vertices. u,v pertencem a V
// A weighted graph is a graph which each edge has a weight
type Edge struct {
	u      int
	v      int
	weight int
}

type Graph map[int][]Edge

func (g Graph) TurnIntoMatrix() [][]int {
	n := len(g)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for u, edges := range g {
		for _, edge := range edges {
			matrix[u-1][edge.v-1] = edge.weight
		}
	}

	return matrix
}

func TurnIntoGraph(n int, edges [][]int) Graph {
	graph := make(Graph)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]

		graph[u] = append(graph[u], Edge{u: u, v: v, weight: w})
		graph[v] = append(graph[v], Edge{u: v, v: u, weight: w})
	}

	return graph
}

func (g Graph) GetNeighbours(node int) []int {
	neighbours := []int{}

	for _, edge := range g[node] {
		neighbours = append(neighbours, edge.v)
	}
	return neighbours
}

func (g Graph) GetDegree(node int) int {
	degree := len(g.GetNeighbours(node))
	return degree
}

type Path []Edge

func countRestrictedPaths(n int, edges [][]int) int {
	fmt.Printf("\nn: %d, edges: %+v", n, edges)
	return -1
}

func distanceToLastNode(x int) int {
	fmt.Println("x: ", x)
	return 0
}
