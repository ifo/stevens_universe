package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.UTC().Format(time.RFC3339))

	newYork, _ := time.LoadLocation("America/New_York")
	fmt.Println(t.In(newYork).Format(time.RFC3339))
}
