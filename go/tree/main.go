package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	n := time.Now()
	tree := Tree{
		Root: &Node{},
	}

	read, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(read)

	input := make(chan string)

	for range make([]struct{}, 100) {
		go tree.Root.Worker(input)
	}

	// only insert one in every 40
	i := 0
	mod := 40
	for scan.Scan() {
		if i%mod == 0 {
			input <- scan.Text()
			i = i / mod
		}
		i++
	}

	log.Println("done")
	log.Println(time.Since(n))
}

func (n *Node) Worker(ch <-chan string) {
	for d := range ch {
		n.AddData([]rune(d), d)
	}
}
