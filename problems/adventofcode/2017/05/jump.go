package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("instructions.txt")
	if err != nil {
		log.Fatalln(err)
	}

	instructionStrings := strings.Split(string(f), "\n")
	insts := Instructions{}
	for _, str := range instructionStrings {
		i, _ := strconv.Atoi(str)
		insts = append(insts, i)
	}

	index := 0
	jumps := 0
	for insts.IsInBounds(index) {
		jumps++
		index = insts.Jump2(index)
	}

	fmt.Printf("Number of jumps: %d\n", jumps)
}

type Instructions []int

func (i Instructions) Jump(index int) int {
	jump := i[index]
	i[index]++
	newIndex := index + jump
	return newIndex
}

func (i Instructions) Jump2(index int) int {
	jump := i[index]
	if jump >= 3 {
		i[index]--
	} else {
		i[index]++
	}
	newIndex := index + jump
	return newIndex
}

func (i Instructions) IsInBounds(index int) bool {
	if index < 0 || index >= len(i) {
		return false
	}
	return true
}
