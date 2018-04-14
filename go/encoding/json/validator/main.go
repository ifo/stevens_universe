package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// This is probably a bad way to validate JSON, but it works.
func main() {
	goodJsonBts := []byte(`{"good":"json"}`)
	badJsonBts := []byte(`{"bad":"json}`)
	var out bytes.Buffer
	err := json.Indent(&out, goodJsonBts, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())

	var out2 bytes.Buffer
	err = json.Indent(&out2, badJsonBts, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}
