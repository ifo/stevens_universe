package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("chain.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var chain Chain
	for _, u := range string(f) {
		chain += Chain(u)
	}

	// Part 1
	fmt.Printf("The length of the reacted chain is: %d\n", chain.ReactLength())

	// Part 2
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	minLength := 0
	for _, r := range alphabet {
		theChain := Chain(
			strings.Replace(
				strings.Replace(string(chain), string(r), "", -1),
				strings.ToUpper(string(r)),
				"",
				-1,
			),
		)
		l := theChain.ReactLength()
		if l < minLength || minLength == 0 {
			minLength = l
		}
	}

	fmt.Printf("The length of the shortest chain is: %d\n", minLength)
}

type Chain string
type Unit rune

func (u Unit) Reacts(u2 Unit) bool {
	if u != u2 && strings.ToLower(string(u)) == strings.ToLower(string(u2)) {
		return true
	}
	return false
}

func (c Chain) ReactLength() int {
	chainLength := len(c)
	chainLengthEnd := 0
	for chainLength != chainLengthEnd {
		chainLength = len(c)
		for i := 0; i+1 < len(c); {
			u := Unit(c[i])
			next := Unit(c[i+1])
			if u.Reacts(next) {
				//fmt.Println(i, i+1)
				c = c[:i] + c[i+2:]
			} else {
				i++
			}
		}
		chainLengthEnd = len(c)
	}
	return chainLength
}
