package handler

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetNote(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	s.mu.Lock()
	defer s.mu.Unlock()

	note, exist := s.notes[id]

	if !exist {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
