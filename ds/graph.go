package ds

import (
	"errors"
)

// Example:
//		0 connects to 2
// 		1 connects to 2 and 3
//		2 connects to 0, 1 and 3
//		3 connects to 1 and 2
//
//					2 -> 0
//				 ^ ^
//				1	-> 3

// Edge List
// An array of connections between the nodes, each inner array lists the connection of node A to node B
//
//	edges = [[0, 2], [2, 3], [2, 1], [1, 3]]

// Adjacent List
// An array/linked list/object that uses an index as its own node's value
// and the actual value is the neighbours' value
//
//	adjacent = [[2], [2,3], [0, 1, 3], [1, 2]]

// Adjacent Matrix
// States 0 and 1 represent if the node has a connection to the another node
// Example:
//
// graph = [
// 	[0, 0, 1, 0],	=>	0 connects to 2
// 	[0, 0, 1, 1],	=>	1 connects to 2 and 3
// 	[1, 1, 0, 1],	=>	2 connects to 0, 1 and 3
// 	[0, 1, 1, 0],	=>	3 connects to 1 and 2
// ]

type Graph struct {
	numOfNodes   int
	adjacentList map[string][]string
}

func (g *Graph) AddVertex(node string) error {
	if g.numOfNodes == 0 {
		g.adjacentList = make(map[string][]string)
	}
	if _, ok := g.adjacentList[node]; ok {
		return errors.New("Vertex/Node already exists.")
	}

	g.adjacentList[node] = make([]string, 0)
	g.numOfNodes++

	return nil
}

func (g *Graph) AddEdge(node1, node2 string) error {
	if _, ok := g.adjacentList[node1]; !ok {
		return errors.New("Vertex/Node \"node1\" does not exist.")
	}
	if _, ok := g.adjacentList[node2]; !ok {
		return errors.New("Vertex/Node \"node2\" does not exist.")
	}

	if !contains(g.adjacentList[node1], node2) {
		g.adjacentList[node1] = append(g.adjacentList[node1], node2)
	}

	if !contains(g.adjacentList[node2], node2) {
		g.adjacentList[node2] = append(g.adjacentList[node2], node1)
	}

	return nil
}

func contains(nodes []string, value string) bool {
	for _, n := range nodes {
		if n == value {
			return true
		}
	}
	return false
}

func (g *Graph) ShowConnections() string {
	result := ""

	for key, nodes := range g.adjacentList {
		connections := ""
		for _, vertex := range nodes {
			connections += vertex + " "
		}
		result += key + "-->" + connections + "\n"
	}
	return result
}
