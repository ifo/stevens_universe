package main

import "fmt"

func main() {
	prog := `package main%simport %q%sfunc main() {%sprog := %s%sn := %q%st := %q%sbt := %q%sfmt.Printf(prog, n+n, "fmt", n+n, n+t, bt+prog+bt, n+t, n, n+t, t, n+t, bt, n+t, n, n)%s}%s`
	n := "\n"
	t := "\t"
	bt := "`"
	fmt.Printf(prog, n+n, "fmt", n+n, n+t, bt+prog+bt, n+t, n, n+t, t, n+t, bt, n+t, n, n)
}
