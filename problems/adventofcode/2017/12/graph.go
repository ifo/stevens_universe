package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("adjacencymatrix.txt")
	if err != nil {
		log.Fatalln(err)
	}

	nodes := strings.Split(string(f), "\n")
	graph := Graph{}
	vGraph := VGraph{}
	for _, n := range nodes {
		node, neighbors := ParseNode(n)
		graph[node] = neighbors
		vGraph[node] = neighbors
	}

	visits := Visits{}
	// Follow all of the nodes starting from node 0
	visits.Follow(0, graph)

	visited := 0
	for _, node := range visits {
		if node {
			visited++
		}
	}

	fmt.Printf("The number of visited nodes: %d\n", visited)

	// Part 2
	visits2 := Visits{}
	groupCount := 0
	for len(vGraph) > 0 {
		for key := range vGraph {
			visits2.FollowDelete(key, vGraph)
			// Only grab the first available key
			break
		}
		groupCount++
	}

	fmt.Printf("The number of groups in the graph: %d\n", groupCount)
}

type Graph [2000][]int
type VGraph map[int][]int

type Visits [2000]bool

func ParseNode(s string) (int, []int) {
	fields := strings.Fields(s)
	node, _ := strconv.Atoi(fields[0])

	adjacents := []int{}
	// Ignore the node name and "<->" fields
	for _, n := range fields[2:] {
		if n[len(n)-1] == ',' {
			n = n[:len(n)-1]
		}
		adj, _ := strconv.Atoi(n)
		adjacents = append(adjacents, adj)
	}

	return node, adjacents
}

func (v *Visits) Follow(node int, g Graph) {
	if v[node] {
		return
	}
	v[node] = true
	for _, n := range g[node] {
		v.Follow(n, g)
	}
}

func (v *Visits) FollowDelete(node int, vg VGraph) {
	if v[node] {
		return
	}
	v[node] = true

	toVisit := vg[node]

	delete(vg, node)
	for _, n := range toVisit {
		v.FollowDelete(n, vg)
	}
}
