package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	templateFiles := map[string]string{
		"define.tmpl": `Something outside the define, shows up only when not calling with "template" {{define "def1"}} Def1 words and context {{.}} {{end}}`,
		"import.tmpl": `Here are some words and context {{.}}. Here is def1: {{template "def1" .Def}}`,
	}

	dir := createTempDir(templateFiles)
	// Ensure that we clean everything up after.
	defer os.RemoveAll(dir)

	// Get all the templates.
	pattern := filepath.Join(dir, "*.tmpl")

	// Parse 'em.
	templates := template.Must(template.ParseGlob(pattern))

	// Print 'em.
	fmt.Println("import.tmpl:")
	templates.ExecuteTemplate(os.Stdout, "import.tmpl", struct{ I, Def string }{I: "import context", Def: "def context"})
	fmt.Println("")
	fmt.Println("define.tmpl:")
	templates.ExecuteTemplate(os.Stdout, "define.tmpl", "define context")
	fmt.Println("")
}

func createTempDir(ts map[string]string) string {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	for name, contents := range ts {
		// Ignore the error to ensure we clean everything up later.
		ioutil.WriteFile(filepath.Join(dir, name), []byte(contents), 0644)
	}

	return dir
}
