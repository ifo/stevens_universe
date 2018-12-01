package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("firewalls.txt")
	if err != nil {
		log.Fatalln(err)
	}

	firewalls := make(Firewalls, 93)
	layers := Layers{}
	firewallInput := strings.Split(string(f), "\n")

	for _, s := range firewallInput {
		fw := ParseFirewall(s)
		firewalls[fw.Depth] = fw
		l := ParseLayer(s)
		layers = append(layers, l)
	}

	// The number of times caught, where an individual severity is depth * range
	severity := 0
	for position := 0; position < len(firewalls); position++ {
		if firewalls.IsCaught(position) {
			severity += firewalls[position].Depth * firewalls[position].Range
		}
		firewalls.Progress()
	}

	fmt.Printf("The severity of the catches is: %d\n", severity)

	// Part 2
	lTestInput := strings.Split("0: 3\n1: 2\n4: 4\n6: 4", "\n")
	lTest := Layers{}
	for _, s := range lTestInput {
		layer := ParseLayer(s)
		lTest = append(lTest, layer)
	}

	testOutput := lTest.FirstSuccess()
	fmt.Println(testOutput)

	firstSuccess := layers.FirstSuccess()
	fmt.Printf("The first successful delay is: %d\n", firstSuccess)
}

type Firewalls []Firewall

type Firewall struct {
	Depth    int
	Range    int
	Position int
	Dir      Direction
}

type Direction int

const (
	Down Direction = iota
	Up
)

func ParseFirewall(s string) Firewall {
	parts := strings.Split(s, ": ")
	depth, _ := strconv.Atoi(parts[0])
	rang, _ := strconv.Atoi(parts[1])

	return Firewall{
		Depth:    depth,
		Range:    rang,
		Position: 0,
		Dir:      Down,
	}
}

func (fws Firewalls) Progress() {
	for i := range fws {
		switch {
		case fws[i].Position == 0:
			fws[i].Dir = Down
			fws[i].Position += 1
		case fws[i].Position == fws[i].Range-1:
			fws[i].Dir = Up
			fws[i].Position -= 1

		case fws[i].Dir == Down:
			fws[i].Position += 1
		case fws[i].Dir == Up:
			fws[i].Position -= 1
		}
	}
}

func (fws Firewalls) IsCaught(position int) bool {
	// Range being 0 means the firewall doesn't actually exist at that spot.
	return fws[position].Position == 0 && fws[position].Range != 0
}

type Layer struct {
	Depth int
	Range int
}

type Layers []Layer

func ParseLayer(s string) Layer {
	parts := strings.Split(s, ": ")
	depth, _ := strconv.Atoi(parts[0])
	rang, _ := strconv.Atoi(parts[1])

	return Layer{
		Depth: depth,
		Range: rang,
	}
}

func (l Layer) Catches(time int) bool {
	return (time+l.Depth)%((l.Range-1)*2) == 0
}

func (l Layers) FirstSuccess() int {
	for i := 0; ; i++ {
		passed := true
		for _, layer := range l {
			if layer.Catches(i) {
				passed = false
				break
			}
		}
		if passed {
			return i
		}
	}
}
