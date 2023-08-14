package data

import (
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	ID        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u UserModel) GetNotes(uid int64) ([]Note, error) {
	query := `SELECT id, user_id
		FROM notes
		WHRE user_id = $1
	`
	rows, err := u.DB.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(
			&note.ID, &note.UserID,
			&note.Content,
			&note.CreatedAt, &note.UpdatedAt,
		)
		if err != nil {
			return notes, err
		}
		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		return notes, err
	}
	return notes, nil
}
