package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("memorybanks.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bankStrings := strings.Split(string(f), "\t")
	b := Bank{}
	for _, str := range bankStrings {
		i, _ := strconv.Atoi(str)
		b = append(b, i)
	}

	iterations := 0
	cycleIterations := 0
	previousStates := map[string]int{}
	for {
		if prevIterations, exists := previousStates[b.ToString()]; exists {
			cycleIterations = iterations - prevIterations
			break
		}
		previousStates[b.ToString()] = iterations
		iterations++
		b.Distribute(b.GetMaxBlocksIndex())
	}

	fmt.Printf("Number of iterations: %d\n", iterations)
	fmt.Printf("Number of iterations in the cycle: %d\n", cycleIterations)
}

type Bank []int

func (b Bank) ToString() string {
	out := ""
	for _, i := range b {
		out += strconv.Itoa(i) + ","
	}
	return out
}

func (b Bank) Distribute(index int) {
	blocks := b[index]
	b[index] = 0
	for ; blocks > 0; blocks-- {
		index++
		b[index%len(b)]++
	}
}

func (b Bank) GetMaxBlocksIndex() int {
	maxBlocksIndex := 0
	for ind := range b {
		if b[maxBlocksIndex] < b[ind] {
			maxBlocksIndex = ind
		}
	}
	return maxBlocksIndex
}
