package data

import (
	"daily-project/internal/validator"
	"database/sql"
	"errors"
	"time"
)

type NoteModel struct {
	DB *sql.DB
}

type Note struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
}

func (n NoteModel) Insert(note *Note) error {
	query := `
		INSERT INTO notes (content)
		VALUES ($1)
		RETURNING id, created_at, updated_at
	`

	// args will hold the $ parameters after VALUES keyword.
	args := []interface{}{note.Content}
	return n.DB.
		QueryRow(query, args...).
		Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt)
}

func (n NoteModel) Get(id int64) (*Note, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, content, created_at, updated_at FROM notes WHERE id = $1`
	var note Note
	err := n.DB.
		QueryRow(query, id).
		Scan(
			&note.ID,
			&note.Content,
			&note.CreatedAt, &note.UpdatedAt,
		)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	// Finalmente retorna el puntero a la nota y error nulo
	return &note, nil
}

func (n NoteModel) Update(note *Note) error {
	query := `UPDATE notes 
		SET content = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING updated_at`

	args := []interface{}{
		note.Content,
		note.ID,
	}

	return n.DB.QueryRow(query, args...).Scan(&note.UpdatedAt)
}

func (n NoteModel) Delete(id int64) error {
	stmt := `DELETE FROM notes WHERE id = $1`
	sqlRes, err := n.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := sqlRes.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func ValidateNote(v *validator.Validator, note *Note) {
	contentSize := len(note.Content)
	v.Check(contentSize > 0 && contentSize <= 255, "content", "can't be empty string or surpass 255 characters")
}
