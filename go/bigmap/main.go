package main

func main() {
	AllPossibleStrings("abcd", 11)
	// wait forever to look at memory usage
	for {
	}
}

// AllPossibleStrings will generate len(alphabet) ^ length number of strings
func AllPossibleStrings(alphabet string, length int) map[string]struct{} {
	out := map[string]struct{}{}
	for _, r := range alphabet {
		out[string(r)] = struct{}{}
	}
	for i := 0; i < length-1; i++ {
		out = AddAnotherCharacter(out, alphabet)
	}
	return out
}

func AddAnotherCharacter(c map[string]struct{}, alphabet string) map[string]struct{} {
	out := map[string]struct{}{}
	for str := range c {
		for _, r := range alphabet {
			out[str+string(r)] = struct{}{}
		}
	}
	return out
}
