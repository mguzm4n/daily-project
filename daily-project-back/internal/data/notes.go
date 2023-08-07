package data

import (
	"daily-project/internal/validator"
	"database/sql"
	"time"
)

type NoteModel struct {
	DB *sql.DB
}

type Note struct {
	ID         int64 `json:"id"`
	CreatedAt  time.Time
	LastEdited time.Time
	Content    string
}

func (n NoteModel) Insert(note *Note) error {
	return nil
}

func (n NoteModel) Get(note *Note) (*Note, error) {
	return nil, nil
}

func (n NoteModel) Update(note *Note) error {
	return nil
}

func (n NoteModel) Delete(id uint64) error {
	return nil
}

func ValidateNote(v *validator.Validator, note *Note) {
	contentSize := len(note.Content)
	v.Check(contentSize > 0 && contentSize <= 255, "content", "can't be empty string or surpass 255 characters")
}
