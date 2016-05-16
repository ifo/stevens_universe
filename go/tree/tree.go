package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Children   [4]*Node
	Data       *Data
	IsComplete bool
	sync.Mutex
}

type Data struct {
	Count  int
	Trials []string
}

func main() {
	tree := Tree{
		Root: &Node{},
	}

	read, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(read)

	// only insert one in every 40
	i := 0
	mod := 40
	for scan.Scan() {
		if i%mod == 0 {
			tree.Root.AddData([]rune(scan.Text()), scan.Text())
			i = i / mod
		}
		i++
	}

	log.Println("done")
	// wait so memory usage can be checked
	for {
	}
}

func (n *Node) AddData(path []rune, data string) {
	// leaf node
	if len(path) == 0 {
		n.Lock()
		if n.Data == nil {
			n.Data = &Data{Count: 1, Trials: []string{data}}
		} else {
			n.Data.Count++
			n.Data.Trials = append(n.Data.Trials, data)
		}
		n.Unlock()
		return
	}

	// recurse to child which exists
	if n.IsComplete {
		n.Children[convertRune(path[0])].AddData(path[1:], data)
		return
	}

	// may need to make child
	ch := convertRune(path[0])
	n.Lock()
	if n.Children[ch] == nil {
		n.Children[ch] = &Node{}
		n.CheckCompleted()
		n.Unlock()
		n.Children[ch].AddData(path[1:], data)
	} else {
		n.Unlock()
		n.Children[ch].AddData(path[1:], data)
	}
}

// CheckCompleted MUST be called from a locked context
func (n *Node) CheckCompleted() {
	for i := range n.Children {
		if n.Children[i] == nil {
			return
		}
	}
	n.IsComplete = true
}

// Traverse returns all leaf nodes
func (n *Node) Traverse() []*Node {
	var out []*Node
	for i := range n.Children {
		if n.Children[i] == nil {
			continue
		}
		if n.Children[i].Data == nil {
			out = append(out, n.Children[i].Traverse()...)
		} else {
			out = append(out, n.Children[i])
		}
	}
	return out
}

func convertRune(r rune) int8 {
	switch r {
	case 'a':
		return 0
	case 'c':
		return 1
	case 'g':
		return 2
	default: // case 't'
		return 3
	}
}

func convertInt8(i int8) rune {
	switch i {
	case 0:
		return 'a'
	case 1:
		return 'c'
	case 2:
		return 'g'
	default: // case 3
		return 't'
	}
}
