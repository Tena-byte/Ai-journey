package repository

import (
	"context"
	"errors"

	"notes-Api/internal/models"

	"github.com/jackc/pgx/v5"
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
func (r *NoteRepository) Create(ctx context.Context, note models.Note) (models.Note, error) {

	query := `
		INSERT INTO notes(title, content)
		VALUES($1, $2)
		RETURNING id, title, content, created_at
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		note.Title,
		note.Content,
	).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.CreatedAt,
	)

	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

func (r *NoteRepository) GetAll(ctx context.Context) ([]models.Note, error) {

	query := `
		SELECT id, title, content, created_at
		FROM notes
		ORDER BY id
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := []models.Note{}

	for rows.Next() {

		var note models.Note

		err := rows.Scan(
			&note.ID,
			&note.Title,
			&note.Content,
			&note.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (r *NoteRepository) GetByID(ctx context.Context, id int) (models.Note, error) {

	query := `
		SELECT id, title, content, created_at
		FROM notes
		WHERE id = $1
	`

	var note models.Note

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Note{}, pgx.ErrNoRows
		}
		return models.Note{}, err
	}

	return note, nil
}

func (r *NoteRepository) Update(ctx context.Context, id int, note models.Note) (models.Note, error) {

	query := `
		UPDATE notes
		SET title = $1,
		    content = $2
		WHERE id = $3
		RETURNING id, title, content, created_at
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		note.Title,
		note.Content,
		id,
	).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Note{}, pgx.ErrNoRows
		}
		return models.Note{}, err
	}

	return note, nil
}

func (r *NoteRepository) Delete(ctx context.Context, id int) error {

	query := `
		DELETE FROM notes
		WHERE id = $1
	`

	result, err := r.DB.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
