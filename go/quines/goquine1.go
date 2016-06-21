package main

import "fmt"

func main() {
	prog := `package main%simport %q%sfunc main() {%sprog := %s%sn := %q%snn := %q%snt := %q%sbt := %q%sfmt.Printf(prog, nn, "fmt", nn, nt, bt+prog+bt, nt, n, nt, nn, nt, nt, nt, bt, nt, n, n)%s}%s`
	n := "\n"
	nn := "\n\n"
	nt := "\n\t"
	bt := "`"
	fmt.Printf(prog, nn, "fmt", nn, nt, bt+prog+bt, nt, n, nt, nn, nt, nt, nt, bt, nt, n, n)
}
