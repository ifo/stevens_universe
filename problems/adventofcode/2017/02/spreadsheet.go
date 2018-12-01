package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("numbers.txt")
	if err != nil {
		log.Fatalln(err)
	}
	inp := string(f)

	parsedSheet := ParseSheet(inp)

	// Part 1
	fmt.Println(SheetChecksum(parsedSheet))

	// Part 2
	fmt.Println(SheetChecksum2(parsedSheet))
}

func ParseSheet(sheet string) [][]int {
	out := [][]int{}
	for _, line := range strings.Split(sheet, "\n") {
		outLine := []int{}
		for _, cell := range strings.Split(line, "\t") {
			n, _ := strconv.Atoi(cell)
			outLine = append(outLine, n)
		}
		out = append(out, outLine)
	}
	return out
}

func SheetChecksum(sheet [][]int) int {
	checksum := 0
	for i := range sheet {
		checksum += RowChecksum(sheet[i])
	}
	return checksum
}

func SheetChecksum2(sheet [][]int) int {
	checksum := 0
	for i := range sheet {
		checksum += RowChecksum2(sheet[i])
	}
	return checksum
}

func RowChecksum(line []int) int {
	large, small := line[0], line[0]
	for _, n := range line[1:] {
		if n > large {
			large = n
		} else if n < small {
			small = n
		}
	}
	return large - small
}

func RowChecksum2(line []int) int {
	sort.Ints(line)
	for i := range line {
		for _, num := range line[i+1:] {
			if num%line[i] == 0 {
				return num / line[i]
			}
		}
	}
	panic("We shouldn't ever get here")
	return -1
}
