package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"github.com/chrismooreproductions/dijkstra/utils/server"
)

// TerminationPoints - the desired start and end point
type TerminationPoints struct {
	start int
	end   int
}

type (
	// Element - a string equivalent of a node
	Element = string
	// Node - An element in the map
	Node = int
	// EdgeList - a list of nodes that can be travelled to from the index
	EdgeList = []Node
	// Graph is an array of edgelists
	Graph = []EdgeList
	// Route - An array of ordered nodes indicating a valid rute through the Graph
	Route = []int
	// FoundRoutes - An array of valid Routes
	FoundRoutes = []Route
)

func main() {
	nodes := [6]Element{"A", "B", "C", "D", "E", "F"}

	//	A = 0
	//	B = 1
	//	C = 2
	//	D = 3
	//	E = 4
	//	F = 5

	//	Demo Graph:
	//	A -> C
	//	B -> C, D, F
	//	C -> A, B, E
	//	D -> A, E
	//	E -> C, D, F
	//	F -> B, E

	var graph = Graph{
		{2},
		{2, 3, 5},
		{0, 1, 4},
		{1, 4},
		{2, 3, 5},
		{1, 4},
	}

	fmt.Printf("Welcome to Dijkstra! Please follow the prompt to get your routes from this list of nodes %v...\n", nodes)

	var tp TerminationPoints
	tp.start = getTerminationPoints(, &nodes)
	tp.end = getTerminationPoints("end", &nodes)

	var foundRoutes, err = run(&graph, tp)
	fmt.Println(foundRoutes, err)
}

func run(graph *Graph, tp TerminationPoints) (FoundRoutes, error) {
	var fr FoundRoutes

	for {
		// Copy initial graph to new array
		g := make(Graph, len(*graph))
		for i := range *graph {
			g[i] = make(EdgeList, len((*graph)[i]))
			copy(g[i], (*graph)[i])
		}

		// Create a Route with start point
		var r Route
		r = append(r, tp.start)

		var match, err = makeRoute(r, g, fr, tp.end)
		if err != nil {
			return fr, errors.New("Unable to find any more routes")
		}
		fr = append(fr, match)
	}
}

func unsetNode(currentNode int, nextNode int, edges Graph) Graph {
	for i, e := range edges {
		if i == currentNode || i == nextNode {
			for j, n := range e {
				if n == nextNode || n == currentNode {
					edges[i] = append(e[:j], e[(j+1):]...)
				}
			}
		}
	}
	return edges
}

func getNextNode(route Route, edges Graph) (int, error) {
	var currentNode = route[len(route)-1]
	if len(edges[currentNode]) > 0 {
		return edges[currentNode][0], nil
	}
	return 0, errors.New("A following node is not available")
}

func backtrack(r Route) Route {
	// Remove last node from route list
	return r[:len(r)-1]
}

func checkExistingMatch(cr Route, fr FoundRoutes) error {
	// If there are no found routes
	if fr == nil {
		return nil
	}

	var isMatch = make([]bool, len(fr))
	for i := range isMatch {
		isMatch[i] = true
	}
	for i := range fr {
		if len(cr) != len(fr[i]) {
			isMatch[i] = false
			continue
		}
		for j, v := range fr[i] {
			if v != cr[j] {
				isMatch[i] = false
				continue
			}
		}
		for j, v := range cr {
			if v != fr[i][j] {
				isMatch[i] = false
				continue
			}
		}
	}
	for _, match := range isMatch {
		if match {
			return errors.New("There was a match")
		}
	}
	return nil
}

func makeRoute(r Route, e Graph, fr FoundRoutes, tpEnd int) (Route, error) {
	// If route length === 0 we must have no possible routes so return error to main()
	if len(r) == 0 {
		return nil, fmt.Errorf("Could not find a route for the given start and end point")
	}

	// The current node
	var currentNode = r[len(r)-1]

	// The next node
	var nextNode, nextNodeErr = getNextNode(r, e)

	e = unsetNode(currentNode, nextNode, e)
	// If there is no further node available from the current last node
	if nextNodeErr != nil {
		// Unset

		r = backtrack(r)
		// and try to find another path...
		return makeRoute(r, e, fr, tpEnd)
	}

	// Otherwise append the nextNode to the route
	r = append(r, nextNode)

	// check if the last node in the route is the end point
	if (r)[len(r)-1] == tpEnd {
		// check for existing match
		var existingMatchErr = checkExistingMatch(r, fr)
		// if route already matches...
		if existingMatchErr != nil {
			// remove the node from the graph
			e = unsetNode(r[len(r)-2], r[len(r)-1], e)
			// back it up
			r = backtrack(r)
			// Try again
			return makeRoute(r, e, fr, tpEnd)
		}
		// else return the route
		return r, nil
	}

	// Try again with the updated route and edges.
	return makeRoute(r, e, fr, tpEnd)
}

func getNodeIndex(nodes *[6]string, node string) (int, error) {
	for i, n := range nodes {
		if n == node {
			return i, nil
		}
	}
	return 0, errors.New("Could not find element ")
}

func getTerminationPoints(position TerminationPoints, nodes *[6]string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your %s point (a single letter from A-F): -> ", position)
	node, _ := reader.ReadString('\n')
	node = strings.Replace(node, "\n", "", -1)
	i, err := getNodeIndex(nodes, node)
	if err != nil {
		fmt.Println(err, node, " Please try again.")
		getTerminationPoints(position, nodes)
	}
	return i
}
