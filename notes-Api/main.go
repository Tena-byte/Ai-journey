package main

import (
	"log"
	"net/http"
	"notes-Api/handler"
	"os"
)

func main() {

	mux := http.NewServeMux()
	server := handler.NewServer()

	
	mux.HandleFunc("/notes", server.Notes)
	

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}
	log.Println("server is running on http://localhost:7000 .....")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Server not running")
	}
}
