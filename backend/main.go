package main

import (
	"log"
	"net/http"
)

func main() {

	// repo  := ""
	// service := ""
	// handler := ""

	log.Println("Server starting on :3500")
	if err := http.ListenAndServe(":3500", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
