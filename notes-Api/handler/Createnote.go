package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) CreateNote(w http.ResponseWriter, r *http.Request) {

	var req CreateNoteRequest

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON request body", 400)
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

	idStr := strconv.Itoa(s.nextID)

	note := Note{
		ID:      idStr,
		Title:   req.Title,
		Content: req.Content,
	}

	s.notes[note.ID] = note
	s.nextID++

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}
