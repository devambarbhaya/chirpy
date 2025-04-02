package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/healthz", HealthzHandler)
	server := &http.Server{
		Handler: mux,
		Addr: ":" + port,
	}
	
	log.Printf("Serving file from %s on port: %s\n", filepathRoot, port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}