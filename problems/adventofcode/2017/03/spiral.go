package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	f, err := ioutil.ReadFile("number.txt")
	if err != nil {
		log.Fatalln(err)
	}
	inputNum, _ := strconv.Atoi(string(f))

	// Random "tests" for clarity.
	fmt.Println(CalculateDistance(10))
	fmt.Println(CalculateDistance(11))
	fmt.Println(CalculateDistance(12))
	fmt.Println(CalculateDistance(13))
	fmt.Println(CalculateDistance(23))
	fmt.Println(CalculateDistance(1024))

	// Part 1
	fmt.Println(CalculateDistance(inputNum))

	// Part 2 - I cheated
	// https://oeis.org/A141481
}

func CalculateDistance(n int) int {
	return GetTier(n) + GetDistanceToMiddle(n) + 1
}

func GetTier(n int) int {
	return GetPreviousSquare(n) / 2
}

func GetPreviousSquare(n int) int {
	for i := range [9223372036854775807 / 2]struct{}{} {
		if i%2 == 0 {
			continue
		}
		if n < i*i {
			return i - 2
		}
	}
	panic("We shouldn't be here")
}

func GetDistanceToMiddle(n int) int {
	t := GetTier(n) + 1
	p := GetPreviousSquare(n)
	index := n - p*p

	normalizedIndex := index % (t * 2)
	distance := t - normalizedIndex
	if distance < 0 {
		distance *= -1
	}

	return distance
}
