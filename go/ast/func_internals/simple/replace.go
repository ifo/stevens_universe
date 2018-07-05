package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

func main() {
	src := `
package main

// main comment
func main() {
	fmt.Println("this is the beginning of the main")
	printHello()
}

// func comment
func printHello() {
	fmt.Print("Hello,")
	fmt.Println(" ast!")
}
`

	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, "src.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Make a map of all func declarations.
	funcMap := map[string]*ast.FuncDecl{}
	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			funcMap[fn.Name.Name] = fn
		}
	}

	// Copy main func statements to a new replacement statement body.
	var newMainBodyList []ast.Stmt
	for _, stmt := range funcMap["main"].Body.List {
		// If this is a function, append its list of statements to the new main list of statements.
		if ident, ok := stmt.(*ast.ExprStmt).X.(*ast.CallExpr).Fun.(*ast.Ident); ok {
			identList := funcMap[ident.Name].Body.List
			newMainBodyList = append(newMainBodyList, identList...)
		} else {
			newMainBodyList = append(newMainBodyList, stmt)
		}
	}

	// Set the main body to be the new list of statements.
	funcMap["main"].Body.List = newMainBodyList

	// Print the modified AST.
	var buf bytes.Buffer
	if err := format.Node(&buf, fileSet, f); err != nil {
		panic(err)
	}
	fmt.Printf("%s", buf.Bytes())
}
