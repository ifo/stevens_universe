package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("boxids.txt")
	if err != nil {
		log.Fatalln(err)
	}

	twoCount := 0
	threeCount := 0
	ids := map[string]struct{}{}
	for _, id := range strings.Split(string(f), "\n") {
		ids[id] = struct{}{}

		characters := map[rune]int{}
		for _, l := range id {
			if _, exists := characters[l]; !exists {
				characters[l] = 1
			} else {
				characters[l]++
			}
		}

		hasTwo := false
		hasThree := false
		for _, count := range characters {
			switch count {
			case 2:
				hasTwo = true
			case 3:
				hasThree = true
			}
		}

		if hasTwo {
			twoCount++
		}
		if hasThree {
			threeCount++
		}
	}

	// Part 1
	fmt.Printf("The checksum of 2 and 3 counts is: %d\n", twoCount*threeCount)

	// Part 2
	outId1 := ""
	outId2 := ""
	for id := range ids {
		zeroMissIds := map[string]struct{}{}
		oneMissIds := map[string]struct{}{}
		for k, v := range ids {
			zeroMissIds[k] = v
		}
		for i, _ := range id {
			for oneMiss, _ := range oneMissIds {
				if oneMiss[i] != id[i] {
					delete(oneMissIds, oneMiss)
				}
			}

			for zeroMiss, _ := range zeroMissIds {
				if zeroMiss[i] != id[i] {
					oneMissIds[zeroMiss] = struct{}{}
					delete(zeroMissIds, zeroMiss)
				}
			}
		}

		if len(oneMissIds) == 1 {
			outId1 = id
			for key := range oneMissIds {
				outId2 = key
			}
			break
		}
	}

	similarities := []byte{}
	for i := range outId1 {
		if outId1[i] == outId2[i] {
			similarities = append(similarities, outId1[i])
		}
	}

	fmt.Printf("The similar id elements are: %s\n", string(similarities))
	// Part 2 faster
	fmt.Printf("The similar id elements are: %s\n", CloseIds(ids))
}

// CloseIds is part2 but O(n*l) instead of O(n^2)
func CloseIds(ids map[string]struct{}) string {
	prefixSuffix := map[string][]string{}
	for id := range ids {
		for i := range id {
			prefix, suffix := id[:i], id[i+1:]
			// We have to check existence, otherwise the default empty string will match
			// the full prefix of the first ID we iterate through.
			if suffixes, exists := prefixSuffix[prefix]; exists {
				for _, suff := range suffixes {
					if suff == suffix {
						return prefix + suffix
					}
				}
			}
			prefixSuffix[prefix] = append(prefixSuffix[prefix], suffix)
		}
	}
	return ""
}
