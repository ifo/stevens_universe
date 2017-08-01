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
	type ctxKey string

	baseContext := context.WithValue(context.Background(), ctxKey("string"), "string")
	baseContext = context.WithValue(baseContext, "string", "string2")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This has a lot of stuff.
		fmt.Printf("%+v\n\n", r.Context())
		// This has nothing but the values we have set.
		fmt.Printf("%+v\n\n", baseContext)
		// This is likely what one wants to use when adding to a request context.
		fmt.Printf("%+v\n\n", context.WithValue(r.Context(), ctxKey("string"), "string"))
		// This erases all context on your request if you pass it through to the next handler.
		fmt.Printf("%+v\n\n", r.WithContext(baseContext).Context())

		// Confirm that ctxKey and string are different things (should print true).
		fmt.Println(baseContext.Value("string").(string) != baseContext.Value(ctxKey("string")).(string))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
