package handler

import (
	"encoding/json"
	"net/http"
)

func (h *NoteHandler) ListNotes(w http.ResponseWriter, r *http.Request) {

	notes, err := h.Repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, "failed to retrieve notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}