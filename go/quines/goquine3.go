package main

import "fmt"

func main() {
	prog := `package main%simport %q%sfunc main() {
	prog := %s%sn := %q%st := %q%sbt := %q%s
	fmt.Printf(prog, n+n, "fmt", n+n, bt+prog+bt,
		n+n+t, n, n+t, t, n+t, bt, n, n, n)%s}%s`

	n := "\n"
	t := "\t"
	bt := "`"

	fmt.Printf(prog, n+n, "fmt", n+n, bt+prog+bt,
		n+n+t, n, n+t, t, n+t, bt, n, n, n)
}
