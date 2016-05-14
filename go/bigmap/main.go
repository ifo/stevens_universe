package main

type Node struct {
	Value string
}

func main() {
	AllPossibleStrings("abcd", 12)
}

// AllPossibleStrings will generate len(alphabet) ^ length number of strings
func AllPossibleStrings(alphabet string, length int) map[string]*Node {
	out := map[string]*Node{}
	for _, r := range alphabet {
		out[string(r)] = &Node{}
	}
	for i := 0; i < length-1; i++ {
		out = AddAnotherCharacter(out, alphabet)
	}
	return out
}

func AddAnotherCharacter(c map[string]*Node, alphabet string) map[string]*Node {
	out := map[string]*Node{}
	for str := range c {
		for _, r := range alphabet {
			out[str+string(r)] = &Node{}
		}
	}
	return out
}
