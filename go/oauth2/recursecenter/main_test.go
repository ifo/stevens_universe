package main

import "testing"

func Test_readToken(t *testing.T) {
	tok, err := readToken("token.json")
	if err != nil {
		t.Error(err)
	}
}
