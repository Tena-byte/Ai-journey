package handler

import (
	"encoding/json"
	"net/http"
	"notes-Api/internal/models"
	"strings"
)

func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {

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

	note := models.Note{
		Title:   req.Title,
		Content: req.Content,
	}

	createdNote, err := h.Repo.Create(r.Context(), note)
	if err != nil {
		http.Error(w, "failed to create note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(createdNote); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
