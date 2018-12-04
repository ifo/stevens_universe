package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("claims.txt")
	if err != nil {
		log.Fatalln(err)
	}

	claims := []Claim{}
	for _, s := range strings.Split(string(f), "\n") {
		c := Claim{}
		fields := strings.Fields(s)

		c.ID, _ = strconv.Atoi(fields[0][1:])

		leftTop := strings.Split(fields[2][:len(fields[2])-1], ",")
		c.Left, _ = strconv.Atoi(leftTop[0])
		c.Top, _ = strconv.Atoi(leftTop[1])

		widthHeight := strings.Split(fields[3], "x")
		c.Width, _ = strconv.Atoi(widthHeight[0])
		c.Height, _ = strconv.Atoi(widthHeight[1])

		claims = append(claims, c)
	}

	fabric := NewFabric(claims)
	for _, c := range claims {
		fabric.Set(c)
	}

	// Part 1
	fmt.Printf("The total overlapping claims are: %d\n", fabric.Count())

	// Part 2
	wholeClaimID := 0
	for _, c := range claims {
		if fabric.CheckClaim(c) {
			wholeClaimID = c.ID
			break
		}
	}

	fmt.Printf("The whole claim ID is: %d\n", wholeClaimID)
}

//type Claim [4]int

type Claim struct {
	ID     int
	Left   int
	Top    int
	Width  int
	Height int
}
type Fabric [][]int

func NewFabric(claims []Claim) Fabric {
	height, width := 0, 0
	for _, claim := range claims {
		rightEdge := claim.Left + claim.Width
		bottomEdge := claim.Top + claim.Height
		if rightEdge > width {
			width = rightEdge
		}
		if bottomEdge > height {
			height = bottomEdge
		}
	}
	out := make(Fabric, height)
	for i := range out {
		out[i] = make([]int, width)
	}
	return out
}

func (f Fabric) Set(c Claim) {
	width := make([]struct{}, c.Width)
	height := make([]struct{}, c.Height)
	for i := range width {
		x := c.Left + i
		for j := range height {
			y := c.Top + j
			f[y][x]++
		}
	}
}

func (f Fabric) Count() int {
	out := 0
	for i := range f {
		for j := range f[i] {
			if f[i][j] > 1 {
				out++
			}
		}
	}
	return out
}

func (f Fabric) CheckClaim(c Claim) bool {
	width := make([]struct{}, c.Width)
	height := make([]struct{}, c.Height)
	for i := range width {
		x := c.Left + i
		for j := range height {
			y := c.Top + j
			if f[y][x] != 1 {
				return false
			}
		}
	}

	return true
}
