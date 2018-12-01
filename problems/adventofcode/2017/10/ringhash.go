package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("lengths.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lengths := []int{}
	for _, l := range strings.Split(string(f), ",") {
		ln, _ := strconv.Atoi(l)
		lengths = append(lengths, ln)
	}

	r := NewRing()
	cursor := 0
	skip := 0
	for _, length := range lengths {
		list := r.Get(cursor, length)
		revList := Reverse(list)
		r.Replace(cursor, revList)
		cursor = (cursor + length + skip) % 256
		skip++
	}

	firstTwoMult := r[0] * r[1]
	fmt.Printf("The multipled first two elements are: %d\n", firstTwoMult)

	// Second part
	// The actual instructions are: Take your same input array (0..255)
	// And use the new lengths example, which is the byte values of the input
	// With `[]byte{17, 31, 73, 47, 23}` at the end.
	blengths := append(f, []byte{17, 31, 73, 47, 23}...)
	br := NewBRing()

	cursor = 0
	skip = 0
	for range [64]struct{}{} {
		for _, l := range blengths {
			length := int(l)
			list := br.Get(cursor, length)
			revList := ReverseBytes(list)
			br.Replace(cursor, revList)
			cursor = (cursor + length + skip) % 256
			skip++
		}
	}

	chunks := make([]byte, 16)
	for i := range chunks {
		parts := br.Get(i*16, 16)
		chunk := byte(parts[0])
		for _, p := range parts[1:] {
			chunk = chunk ^ byte(p)
		}
		chunks[i] = chunk
	}

	fmt.Println(hex.EncodeToString(chunks))
}

type Ring [256]int
type BRing [256]byte

func NewRing() Ring {
	r := Ring{}
	for i := range r {
		r[i] = i
	}
	return r
}

func (r *Ring) Replace(cursor int, list []int) {
	for i := 0; i < len(list); i++ {
		index := (i + cursor) % 256
		r[index] = list[i]
	}
}

func (r *Ring) Get(cursor, length int) []int {
	list := []int{}
	for i := cursor; i < cursor+length; i++ {
		list = append(list, r[i%256])
	}
	return list
}

func Reverse(l []int) []int {
	out := []int{}
	for _, v := range l {
		out = append([]int{v}, out...)
	}
	return out
}

func NewBRing() BRing {
	r := BRing{}
	for i := range r {
		r[i] = byte(i)
	}
	return r
}

func (r *BRing) Replace(cursor int, list []byte) {
	for i := 0; i < len(list); i++ {
		index := (i + cursor) % 256
		r[index] = list[i]
	}
}

func (r *BRing) Get(cursor, length int) []byte {
	list := []byte{}
	for i := cursor; i < cursor+length; i++ {
		list = append(list, r[i%256])
	}
	return list
}

func ReverseBytes(l []byte) []byte {
	out := []byte{}
	for _, v := range l {
		out = append([]byte{v}, out...)
	}
	return out
}
