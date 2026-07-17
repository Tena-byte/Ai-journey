package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"notes-Api/internal/models"

	"github.com/jackc/pgx/v5"
)

func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid note ID", http.StatusBadRequest)
		return
	}

	var req CreateNoteRequest

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		http.Error(w, "title cannot be empty", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Content) == "" {
		http.Error(w, "content cannot be empty", http.StatusBadRequest)
		return
	}

	note := models.Note{
		Title:   req.Title,
		Content: req.Content,
	}

	updatedNote, err := h.Repo.Update(r.Context(), id, note)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "note not found", http.StatusNotFound)
			return
		}

		http.Error(w, "failed to update note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(updatedNote); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
