package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	f, err := ioutil.ReadFile("sequence.txt")
	if err != nil {
		log.Fatalln(err)
	}
	inp := string(f)

	// Part 1
	fmt.Println(Captcha(inp))

	// Part 2
	fmt.Println(Captcha2(inp))
}

func Captcha(nums string) int {
	out := 0
	for i := range nums {
		if nums[i] == nums[(i+1)%len(nums)] {
			n, _ := strconv.Atoi(string(nums[i]))
			out += n
		}
	}
	return out
}

func Captcha2(nums string) int {
	out := 0
	offset := len(nums) / 2
	for i := range nums {
		if nums[i] == nums[(i+offset)%len(nums)] {
			n, _ := strconv.Atoi(string(nums[i]))
			out += n
		}
	}
	return out
}
