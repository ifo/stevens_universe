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

	// only insert one in every 40
	i := 0
	mod := 40
	for scan.Scan() {
		if i%mod == 0 {
			text := scan.Text()
			go func() {
				tree.Root.AddData([]rune(text), text)
			}()
			i = i / mod
		}
		i++
	}

	log.Println("done")
	log.Println(time.Since(n))
}
