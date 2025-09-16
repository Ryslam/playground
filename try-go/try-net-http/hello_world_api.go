package main

import (
	"fmt"
	"net/http"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", rootHandler)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
