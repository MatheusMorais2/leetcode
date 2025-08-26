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
	fmt.Print("neighbours: \n", neighbours)
	degree := graph.GetDegree(4)
	fmt.Print("degree: \n", degree)
	//countRestrictedPaths(5, edges)
	matrix := graph.TurnIntoMatrix()
	for i := 0; i < len(matrix); i++ {
		fmt.Printf("%d: ", i+1)
		fmt.Printf(" %d \n", matrix[i])
	}
}

// Graph G = (V, E) is a set of vertices V and edges E,
// where each edge (u,v) is a conection between vertices. u,v pertencem a V
type Edge struct {
	u      int
	v      int
	weight int
}

type Graph map[int][]Edge

// WRONG!!
func (g Graph) TurnIntoMatrix() [][]int {
	matrix := [][]int{}
	for i := 0; i < len(g); i++ {
		node := g[i]
		matrix = append(matrix, []int{})
		fmt.Printf("\ni=%d, matrix: %+v", i, matrix)
		for _, edge := range node {
			if edge.weight != 0 {
				matrix[i] = append(matrix[i], edge.weight)
			} else {
				matrix[i] = append(matrix[i], 0)
			}
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

func DepthFirstSearch(initialNode int, finalNode int, nodes Graph) {
}

func distanceToLastNode(x int) int {
	fmt.Println("x: ", x)
	return 0
}
