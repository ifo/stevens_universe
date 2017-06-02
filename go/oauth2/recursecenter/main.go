// Blatant copy of https://godoc.org/golang.org/x/oauth2#example-Config
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ifo/oauth2rc"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Endpoint:     oauth2rc.Endpoint,
		//Scopes:       []string{""}, // not necessary in this case
	}

	var client *http.Client
	// See if a usable token exists.
	tok, err := readToken("token.json")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// Not checking the error because it doesn't matter here.
	if tok == nil {
		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		// P.S. You should randomize the string you send in the first argument here,
		// and check it when it comes back, to ensure it's the same one you sent.
		fmt.Printf("Visit the URL for the auth dialog: %v", url)

		// Use the authorization code that is pushed to the redirect
		// URL. Exchange will do the handshake to retrieve the
		// initial access token. The HTTP Client returned by
		// conf.Client will refresh the token as necessary.
		var code string
		if _, err := fmt.Scan(&code); err != nil {
			log.Fatal(err)
		}
		tok, err = conf.Exchange(ctx, code)
		if err != nil {
			log.Fatal(err)
		}

		err = writeToken(tok, "token.json")
		if err != nil {
			log.Fatal(err)
		}

		client = conf.Client(ctx, tok)
	} else {
		client = oauth2.NewClient(ctx, conf.TokenSource(ctx, tok))
	}

	resp, err := client.Get("https://www.recurse.com/api/v1/people/me")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Pretty print the response.
	var jsonBody bytes.Buffer
	err = json.Indent(&jsonBody, body, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// You should see some info about yourself.
	fmt.Println(jsonBody.String())
}

// Check for token.
func readToken(file string) (*oauth2.Token, error) {
	tokBts, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var t oauth2.Token
	err = json.Unmarshal(tokBts, &t)
	return &t, err
}

// Save token.
func writeToken(t *oauth2.Token, file string) error {
	tokenJson, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, tokenJson, 0644)
}
