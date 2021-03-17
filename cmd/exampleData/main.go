package main

import (
	"bytes"
	_ "embed"
	"log"
	"net/http"
)

//go:embed books.json
var exampleData []byte

func main() {

	// Send the example JSON to the server.
	_, err := http.Post("http://bookstore.micahparks.com/api/books/upsert", "application/json", bytes.NewReader(exampleData))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
