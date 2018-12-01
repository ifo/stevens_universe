package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("changes.txt")
	if err != nil {
		log.Fatalln(err)
	}

	frequency := 0

	changeInput := strings.Split(string(f), "\n")
	changes := []int{}
	for _, s := range changeInput {
		direction := 1
		if s[0] == '-' {
			direction = -1
		}

		change, _ := strconv.Atoi(s[1:])
		changes = append(changes, change*direction)
		frequency += change * direction
	}

	// Part 1
	fmt.Printf("The total frequency is: %d\n", frequency)

	// Part 2
	repeatedFreq := FirstRepeatedFrequency(changes)
	fmt.Printf("The first repeated frequency is: %d\n", repeatedFreq)
}

func FirstRepeatedFrequency(changes []int) int {
	freq := 0
	seenFreqs := map[int]struct{}{freq: struct{}{}}
	for {
		for _, c := range changes {
			freq += c
			if _, exists := seenFreqs[freq]; exists {
				return freq
			}
			seenFreqs[freq] = struct{}{}
		}
	}
}
