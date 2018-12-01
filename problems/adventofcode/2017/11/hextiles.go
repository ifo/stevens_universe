package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("movements.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//movements := map[string]int{"n": 0, "s": 0, "ne": 0, "nw": 0, "se": 0, "sw": 0}
	movements := Movements{}
	maxDistance := 0
	for _, m := range strings.Split(string(f), ",") {
		switch m {
		case "n":
			movements[N]++
		case "ne":
			movements[NE]++
		case "se":
			movements[SE]++
		case "s":
			movements[S]++
		case "sw":
			movements[SW]++
		case "nw":
			movements[NW]++
		}
		if d := movements.Distance(); d > maxDistance {
			maxDistance = d
		}
	}

	endDistance := movements.Distance()

	// Part 1
	fmt.Printf("The ending distance is: %d\n", endDistance)

	// Part 2
	fmt.Printf("The furthest distance was: %d\n", maxDistance)
}

type Direction int

type Movements [6]int

const (
	N Direction = iota
	NE
	SE
	S
	SW
	NW
)

func (m Movements) Distance() int {
	directions := []int{}
	for dir := range m {
		distance := m[dir] - m[(dir+3)%6]
		if distance > 0 {
			directions = append(directions, distance)
		}
	}

	if len(directions) == 1 {
		return directions[0]
	}
	if len(directions) == 2 {
		return directions[0] + directions[1]
	}

	distance := directions[1]
	// The first and last of the 3 elements in directions can cancel into one
	// of the middle direction.
	if directions[0] > directions[2] {
		distance += directions[0]
	} else {
		distance += directions[2]
	}

	return distance
}
