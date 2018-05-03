package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const port = ":8080"

type user struct {
	Name  string
	Email string
}

// function handler
func handler(w http.ResponseWriter, r *http.Request) {
	u := user{Name: "Jogi", Email: "Silalahi"}

	json, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	// Route
	http.HandleFunc("/", handler)

	// Start the server with specidied port
	log.Fatal(http.ListenAndServe(port, nil))
}
