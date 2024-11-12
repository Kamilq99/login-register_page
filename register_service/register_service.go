package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Register Service start on localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
