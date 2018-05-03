package main

import "github.com/jsilalahi/learn-go/import/db"

func main() {
	// This is error, since the function is private
	// db.open()

	db.Close()
}
