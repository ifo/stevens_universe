package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("stream.txt")
	if err != nil {
		log.Fatalln(err)
	}

	state := Group
	nestLevel := 0
	total := 0
	garbageCharacters := 0
	for _, r := range strings.Split(string(f), "") {
		if state == IgnoreNext {
			state = Garbage
			continue
		}
		if state == Garbage {
			switch r {
			case "!":
				state = IgnoreNext
			case ">":
				state = Group
			default:
				garbageCharacters++
			}
		}
		if state == Group {
			switch r {
			case "<":
				state = Garbage
			case "{":
				nestLevel++
				total += nestLevel
			case "}":
				nestLevel--
			}
		}
	}

	fmt.Printf("Value of groups: %d\n", total)
	fmt.Printf("Number of garbage characters: %d\n", garbageCharacters)
}

type State int

const (
	Group State = iota
	Garbage
	IgnoreNext
)
