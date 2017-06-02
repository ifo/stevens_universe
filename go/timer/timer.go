package main

import (
	"fmt"
	"time"
)

func main() {
	for i := range [4]struct{}{} {
		time.Sleep(5 * time.Second)
		fmt.Println((i + 1) * 5)
	}
	fmt.Println("done")
}
