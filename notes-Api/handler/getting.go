package handler

import (
	"encoding/json"
	"net/http"
)





func (s *Server) ListNotes(w http.ResponseWriter, r *http.Request){
	
	s.mu.Lock()
	defer s.mu.Unlock()

	noteList := []Note{}

	for _, nts := range s.notes{

		noteList = append(noteList, nts)
	}

	w.Header().Set("Content-Type", "application/json")
	
	if err := json.NewEncoder(w).Encode(noteList); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}