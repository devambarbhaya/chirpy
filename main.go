package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/", fileServer)

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