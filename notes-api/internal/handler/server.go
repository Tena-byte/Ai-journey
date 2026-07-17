package handler

import "notes-Api/internal/repository"

type NoteHandler struct {
	Repo *repository.NoteRepository
}

func NewNoteHandler(repo *repository.NoteRepository) *NoteHandler {
	return &NoteHandler{
		Repo: repo,
	}
}