package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

// function handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	// Route
	http.HandleFunc("/", handler)

	// Start the server with specidied port
	log.Fatal(http.ListenAndServe(port, nil))
}
