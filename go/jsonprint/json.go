package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	jsonBytes := []byte(`{"one":1,"two":[1,2]}`)

	var bPrint bytes.Buffer
	err := json.Indent(&bPrint, jsonBytes, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	obj := Obj{
		One: "one",
		Two: []int{1, 2},
	}

	objPrint, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Indent Bytes (two spaces):\n%s\n", bPrint.String())
	fmt.Printf("MarshalIndent Struct (tabs):\n%s\n", string(objPrint))
}

type Obj struct {
	One string `json:"one"`
	Two []int  `json:"two"`
}
