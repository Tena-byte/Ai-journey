package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) CreateNote(w http.ResponseWriter, r *http.Request) {

	var req CreateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON request body", 400)
		return
	}

	defer r.Body.Close()

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

	idStr := strconv.Itoa(len(s.notes) + 1)

	note := Note{
		ID:      idStr,
		Title:   req.Title,
		Content: req.Content,
	}

	s.notes[note.ID] = note
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}
