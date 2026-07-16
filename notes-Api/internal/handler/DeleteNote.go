package handler

import (
	"encoding/json"
	"net/http"
)

func (s *Server) DeleteNote(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.notes[id]

	if !ok {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	message := "Note deleted successfully"

	delete(s.notes, id)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, "failded to encode response", http.StatusInternalServerError)
		return
	}
}
