package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// It is recommended to not use builtin types for context keys, to prevent overwriting
	// other people's changes.
	// Making a simple string type wrapper prevents context key collision.
	// NOTE: It prevents collision even if another package uses a type with the same name,
	// as the same named type from different packages are still different types.
	type contextKey string

	baseContext := context.WithValue(context.Background(), contextKey("string"), "string")
	baseContext = context.WithValue(baseContext, "string", "string2")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Here are some examples of contexts.\n\n"))

		// This has a lot of stuff.
		w.Write([]byte("r.Context():\n"))
		w.Write([]byte(fmt.Sprintf("%+v\n\n", r.Context())))
		// This has nothing but the values we have set.
		w.Write([]byte("baseContext:\n"))
		w.Write([]byte(fmt.Sprintf("%+v\n\n", baseContext)))
		// This is likely what one wants to use when adding to a request context.
		w.Write([]byte("context.WithValue(r.Context(), contextKey(\"string\"), \"string\"):\n"))
		w.Write([]byte(
			fmt.Sprintf("%+v\n\n", context.WithValue(r.Context(), contextKey("string"), "string"))))
		// This erases all context on your request if you pass it through to the next handler.
		w.Write([]byte("r.WithContext(baseContext).Context():\n"))
		w.Write([]byte(fmt.Sprintf("%+v\n\n", r.WithContext(baseContext).Context())))

		// Confirm that contextKey and string are different things (should print true).
		w.Write([]byte(
			"baseContext.Value(\"string\").(string) != baseContext.Value(contextKey(\"string\")).(string):\n"))
		w.Write([]byte(fmt.Sprintln(
			baseContext.Value("string").(string) != baseContext.Value(contextKey("string")).(string))))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
