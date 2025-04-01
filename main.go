package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr: ":8080",
	}
	log.Println("Server starting on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}