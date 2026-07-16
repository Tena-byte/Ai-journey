package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (s *Server) UpdateNote(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var req CreateNoteRequest

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		http.Error(w, "Title can not be empty", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Content) == "" {
		http.Error(w, "Content can not be empty", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	note, exists := s.notes[id]

	if !exists {
		http.Error(w, "Id Not Found", http.StatusNotFound)
		return
	}

	note.Title = req.Title
	note.Content = req.Content
	s.notes[id] = note

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}
