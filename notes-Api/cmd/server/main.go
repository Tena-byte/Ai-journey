package main

import (
	"log"
	"net/http"
	"notes-Api/internals/handler"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	mux := http.NewServeMux()
	server := handler.NewServer()

	mux.HandleFunc("POST /notes", server.CreateNote)
	mux.HandleFunc("GET /notes", server.ListNotes)

	mux.HandleFunc("GET /notes/{id}", server.GetNote)
	mux.HandleFunc("PUT /notes/{id}", server.UpdateNote)
	mux.HandleFunc("DELETE /notes/{id}", server.DeleteNote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}
	log.Printf("server is running on http://localhost:%s ...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
