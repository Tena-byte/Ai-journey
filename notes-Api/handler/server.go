package handler

import (
	"sync"
)

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}


type Server struct {
	notes map[string]Note
	mu    sync.Mutex
}

func NewServer() *Server {
	return &Server{
		notes: make(map[string]Note),
	}
}
