package handler

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
)

func (s *Server) ListNotes(w http.ResponseWriter, r *http.Request) {

	s.mu.Lock()
	defer s.mu.Unlock()

	noteList := []Note{}
	keys := []int{}

	for key := range s.notes {
		id, err := strconv.Atoi(key)
		if err != nil {
			http.Error(w, "invalid note ID", http.StatusInternalServerError)
			return
		}
		keys = append(keys, id)

	}

	slices.Sort(keys)

	for _, id := range keys {
		key := strconv.Itoa(id)
		note := s.notes[key]
		noteList = append(noteList, note)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(noteList); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
