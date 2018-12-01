package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("passphrases.txt")
	if err != nil {
		log.Fatalln(err)
	}

	passphrases := strings.Split(string(f), "\n")

	validPassphrases := 0
	for _, p := range passphrases {
		if validatePassphrase2(p) {
			validPassphrases++
		}
	}

	fmt.Printf("number of valid passphrases: %d\n", validPassphrases)
}

func validatePassphrase(s string) bool {
	words := map[string]struct{}{}
	for _, word := range strings.Fields(s) {
		if _, exists := words[word]; exists {
			return false
		}
		words[word] = struct{}{}
	}
	return true
}

func validatePassphrase2(s string) bool {
	words := map[string]struct{}{}
	for _, word := range strings.Fields(s) {
		sortedWord := sortString(word)
		if _, exists := words[sortedWord]; exists {
			return false
		}
		words[sortedWord] = struct{}{}
	}
	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
