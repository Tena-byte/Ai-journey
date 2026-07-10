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

	mux.HandleFunc("GET /notes", server.ListNotes)
	mux.HandleFunc("POST /notes", server.CreateNote)

	//mux.HandleFunc("GET /notes/{id}", server.GetNote)
	mux.HandleFunc("PUT /notes/{id}", server.UpdateNote)
	//mux.HandleFunc("DELETE /notes/{id}", server.DeleteNote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}
	log.Println("server is running on http://localhost:7000 .....")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Server not running")
	}
}
