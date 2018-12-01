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

	treeNodes := strings.Split(string(f), "\n")
	// Get all nodes and note all children
	allNodesByName := map[string]Node{}
	allChildNodeNames := map[string]struct{}{}
	for _, str := range treeNodes {
		node := ParseNode(str)
		allNodesByName[node.Name] = node
		for _, childNode := range node.Children {
			nodeName := childNode.Name
			allChildNodeNames[nodeName] = struct{}{}
		}
	}

	// Part 1
	parentNodeName := ""
	for _, n := range allNodesByName {
		if _, exists := allChildNodeNames[n.Name]; !exists {
			parentNodeName = n.Name
			break
		}
	}
	fmt.Printf("Parent node name: %s\n", parentNodeName)

	// Part 2
	// Properly fill in children
	tree := allNodesByName[parentNodeName]
	tree.CalculateWeights(allNodesByName)
	node, weight := tree.FindAnomoly(0)
	properWeight := node.GetProperWeight(weight)
	fmt.Printf("The proper weight is: %d\n", properWeight)
}

type Node struct {
	Name          string
	Weight        int
	Children      []Node
	SubTreeWeight int
}

func (n *Node) CalculateWeights(nodes map[string]Node) {
	// Our end case is having calculated the subtree weight
	if n.SubTreeWeight != 0 {
		return
	}

	childrenTotals := 0
	for i := range n.Children {
		name := n.Children[i].Name
		// Hydrate child node
		n.Children[i] = nodes[name]
		// Calculate subtree weight
		n.Children[i].CalculateWeights(nodes)
		childrenTotals += n.Children[i].SubTreeWeight
	}

	n.SubTreeWeight = n.Weight + childrenTotals
}

// FindAnomoly returns the node that is the anomoly and the proper weight
func (n Node) FindAnomoly(properWeight int) (Node, int) {
	hasAnomoly := false
	anomolyNode := n.Children[0]
	normalWeight := n.Children[0].SubTreeWeight
	for _, c := range n.Children[1:] {
		if c.SubTreeWeight != normalWeight {
			// If we're here twice, the first node is the anomoly
			if hasAnomoly == true {
				anomolyNode = n.Children[0]
				normalWeight = c.SubTreeWeight
				break
			}
			hasAnomoly = true
			anomolyNode = c
		}
	}

	if hasAnomoly {
		return anomolyNode.FindAnomoly(normalWeight)
	}
	return n, properWeight
}

func (n Node) GetProperWeight(weight int) int {
	diff := n.SubTreeWeight - weight
	if diff < 0 {
		diff *= -1
	}
	return n.Weight - diff
}

func ParseNode(s string) Node {
	fields := strings.Fields(s)
	name := fields[0]
	// Parse the weight number after removing the surrounding parens: "()"
	weight, _ := strconv.Atoi(fields[1][1 : len(fields[1])-1])

	node := Node{Name: name, Weight: weight}
	if len(fields) < 3 {
		// Nodes without children have a subtree weight equal to their own weight
		node.SubTreeWeight = weight
		return node
	}

	// Skip past the "->"
	for _, name := range fields[3:] {
		if name[len(name)-1] == ',' {
			name = name[:len(name)-1]
		}
		node.Children = append(node.Children, Node{Name: name})
	}
	return node
}
