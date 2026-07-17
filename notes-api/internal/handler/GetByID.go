package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func (h *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid note ID", http.StatusBadRequest)
		return
	}

	note, err := h.Repo.GetByID(r.Context(), id)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "note not found", http.StatusNotFound)
			return
		}

		http.Error(w, "failed to retrieve note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
