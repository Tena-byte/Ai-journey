package handler

import (
	"fmt"
	"net/http"
	"sync"
)

type Note struct {
	ID      string `json:"id"`
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

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {

	l := len(s.notes)
	fmt.Fprintln(w, "Hello World", l)

	fmt.Fprintf(w, "Server address: %p\n", s)
}
