package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type node struct {
	value    int
	children []*node
}

func (n *node) DepthFirstSearch() (dfs []int) {
	dfs = append(dfs, n.value)

	for _, child := range n.children {
		dfs = append(dfs, child.DepthFirstSearch()...)
	}

	return dfs
}

func (n *node) BreadthFirstSearch() (bfs []int) {
	queue := []*node{n}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		bfs = append(bfs, curr.value)

		for _, child := range curr.children {
			queue = append(queue, child)
		}
	}

	return bfs
}

func (n *node) getBreadth() int {
	w := 1

	queue := []*node{n}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.children != nil {
			w += len(curr.children) - 1
		}

		for _, child := range curr.children {
			queue = append(queue, child)
		}
	}

	return w
}

func (n *node) getDepth() int {
	d := make([]int, len(n.children))

	for i, child := range n.children {
		d[i] = child.getDepth()
	}

	var maxD int

	for _, el := range d {
		if maxD < el {
			maxD = el
		}
	}

	return maxD + 1
}

func main() {
	// fN, gV := readGraph() // to input from CL

	fN := 1 // first (head) node value
	gV := make(map[int][]int)
	gV[1] = []int{2, 3, 4} // map with node value key and children elements
	gV[2] = []int{5, 6}
	gV[5] = []int{9, 10}
	gV[4] = []int{7, 8} // manual graph input (preferred)

	graph := mapToGraph(fN, gV)

	fmt.Printf("Depth: %d\n", graph.getDepth())
	fmt.Printf("Breadth: %d\n", graph.getBreadth())

	t := time.Now()
	dfs := graph.DepthFirstSearch()
	tookDFS := time.Since(t)

	t = time.Now()
	bfs := graph.BreadthFirstSearch()
	tookBFS := time.Since(t)

	fmt.Printf("DFS took %v, res: %v\n", tookDFS, dfs)
	fmt.Printf("BFS took %v, res: %v\n", tookBFS, bfs)

	switch {
	case tookDFS > tookBFS:
		fmt.Println("BFS is more effective with a given graph")
	case tookDFS < tookBFS:
		fmt.Println("DFS is more effective with a given graph")
	default:
		fmt.Println("Looks like both methods are equally good with this graph")
	}
}

func mapToGraph(fN int, m map[int][]int) node {
	var children []*node

	if _, ok := m[fN]; !ok {
		return node{
			value:    fN,
			children: nil,
		}
	}

	for _, val := range m[fN] {
		child := mapToGraph(val, m)

		children = append(children, &child)
	}

	return node{
		value:    fN,
		children: children,
	}
}

func readGraph() (int, map[int][]int) {
	var (
		firstNode  int
		firstFound bool
	)

	graphValues := make(map[int][]int)

	fmt.Println("Enter a node followed by its children separated by spaces, row by row (values must be unique)")

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		strElems := strings.Split(line, " ")
		val, _ := strconv.Atoi(strElems[0])

		if !firstFound {
			firstFound = true
			firstNode = val
		}

		chldrn := make([]int, len(strElems)-1)

		for i := 1; i < len(strElems); i++ {
			chldrn[i-1], _ = strconv.Atoi(strElems[i])
		}

		graphValues[val] = chldrn
	}

	return firstNode, graphValues
}
