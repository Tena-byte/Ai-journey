package repository

import (
	"context"

	
	"notes-Api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository struct {
	DB *pgxpool.Pool
}

func NewNoteRepository(db *pgxpool.Pool) *NoteRepository {
	return &NoteRepository{
		DB: db,
	}
}

func (r *NoteRepository) Create(
	ctx context.Context,
	note models.Note,
) error {

	query := `
		INSERT INTO notes(title, content)
		VALUES($1, $2)
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		note.Title,
		note.Content,
	)

	return err
}
