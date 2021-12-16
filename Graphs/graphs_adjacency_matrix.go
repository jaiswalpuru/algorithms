package main

import (
	"errors"
	"fmt"
)

type GraphType string

const (
	DIRECTED   GraphType = "DIRECTED"
	UNDIRECTED GraphType = "UNDIRECTED"
)

type Graph interface {
	Init()
	AddEdge(vertexOne, vertexTwo int) error
	AddEdgeWithWeight(vertexOne, vertexTwo, weight int) error
	RemoveEdge(vertexOne, vertexTwo int) error
	HasEdge(vertexOne, vertexTwo int) bool
	GetGraphType() GraphType
	GetAdjacencyNodesForVertex(vertex int) map[int]bool
	GetWeightOfEdge(vertexOne, vertexTwo int) (int, error)
	GetNumberOfVertices() int
	GetNumberOfEdges() int
	GetInDegreeOfVertex(vertex int) int
}

type AdjacencyMatrix struct {
	Vertices  int
	Edges     int
	GraphType GraphType
	AdjMatrx  [][]int
}

func (G *AdjacencyMatrix) Init() {
	G.AdjMatrx = make([][]int, G.Vertices)
	G.Edges = 0
	for i := 0; i < G.Vertices; i++ {
		G.AdjMatrx[i] = make([]int, G.Vertices)
	}
}

func (G *AdjacencyMatrix) AddEdge(vertexOne, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	G.AdjMatrx[vertexOne][vertexTwo] = 1
	G.Edges++
	if G.GraphType == UNDIRECTED {
		G.AdjMatrx[vertexTwo][vertexOne] = 1
		G.Edges++
	}
	return nil
}

func (G *AdjacencyMatrix) AddEdgeWithWeight(vertexOne, vertexTwo, weight int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	G.AdjMatrx[vertexOne][vertexTwo] = weight
	G.Edges++
	if G.GraphType == UNDIRECTED {
		G.AdjMatrx[vertexTwo][vertexOne] = weight
		G.Edges++
	}
	return nil
}

func (G *AdjacencyMatrix) RemoveEdge(vertexOne, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Imdex out of bounds")
	}
	G.AdjMatrx[vertexOne][vertexTwo] = 0
	G.Edges--
	if G.GraphType == UNDIRECTED {
		G.AdjMatrx[vertexTwo][vertexOne] = 0
		G.Edges--
	}
	return nil
}

func (G *AdjacencyMatrix) HasEdge(vertexOne, vertexTwo int) bool {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return false
	}
	return G.AdjMatrx[vertexOne][vertexTwo] != 0
}

func (G *AdjacencyMatrix) GetGraphType() GraphType {
	return G.GraphType
}

func (G *AdjacencyMatrix) GetAdjacencyNodesForVertex(vertex int) map[int]bool {
	adjancencyMatrucVertices := map[int]bool{}

	if vertex >= G.Vertices || vertex < 0 {
		return adjancencyMatrucVertices
	}
	for i := 0; i < G.Vertices; i++ {
		if G.AdjMatrx[vertex][i] != 0 {
			adjancencyMatrucVertices[i] = (G.AdjMatrx[vertex][i] != 0)
		}
	}
	return adjancencyMatrucVertices
}

func (G *AdjacencyMatrix) GetWeightOfEdge(vertexOne, vertexTwo int) (int, error) {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return 0, errors.New("Error getting weight for the vertex")
	}
	return G.AdjMatrx[vertexOne][vertexTwo], nil
}

func (G *AdjacencyMatrix) GetNumberOfVertices() int {
	return G.Vertices
}

func (G *AdjacencyMatrix) GetNumberOfEdges() int {
	return G.Edges
}

func (G *AdjacencyMatrix) GetInDegreeOgVertex(vertex int) int {
	inDegree := 0
	adjacencyNodes := G.GetAdjacencyNodesForVertex(vertex)
	for key := range adjacencyNodes {
		if adjacencyNodes[key] {
			inDegree++
		}
	}
	return inDegree
}

func main() {
	var testAdjMatrixDirected = &AdjacencyMatrix{4, 0, DIRECTED, nil}
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdge(2, 1)
	if err != nil {
		fmt.Printf("Error adding edge")
	}
	if testAdjMatrixDirected.AdjMatrx[2][1] != 1 {
		fmt.Println("Data not found at index")
	}
	if testAdjMatrixDirected.AdjMatrx[1][2] != 0 {
		fmt.Printf("Data not found at index")
	}
}
