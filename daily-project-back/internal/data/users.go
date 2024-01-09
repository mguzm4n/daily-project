package data

import (
	"context"
	"daily-project/internal/validator"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type UserModel struct {
	DB *sql.DB
}

type password struct {
	plaintext *string
	hash      []byte
}

type User struct {
	ID        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
}

func (p *password) Set(plaintextPwd string) error {
	cost := 12

	// The higher the cost, function becomes slower computationally.
	// Returns a hash with a salt prefix.
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPwd), cost)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPwd
	p.hash = hash

	return nil
}

// Checks if string password matches the hash result already stored.
func (p *password) Matches(plaintextPwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPwd))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(
		validator.Matches(email, validator.EmailRX),
		"email",
		"must be a valid email address",
	)
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Firstname != "", "name", "must be provided")
	v.Check(len(user.Firstname) <= 500, "name", "must not be more than 500 bytes long")

	// Call the standalone ValidateEmail() helper.
	ValidateEmail(v, user.Email)

	// If the plaintext password is not nil, call the standalone
	// ValidatePasswordPlaintext() helper.
	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	// If the password hash is ever nil, this will be due to a logic
	// error in our codebase (probably because we forgot to set a
	// password for the user). It's a useful sanity check to include
	// here, but it's not a problem with the data provided by the client.
	// So rather than adding an error to the validation map we
	// raise a panic instead.
	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
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

func (m UserModel) Insert(user *User) error {
	query := `INSERT INTO users(firstname, lastname, email, password_hash) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	args := []interface{}{user.Firstname, user.Lastname, user.Email, user.Password.hash}
	err := m.DB.QueryRow(query, args...).
		Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

// Retrieve the User details from the database based on the user's email address.
// Because we have a UNIQUE constraint on the email column, this SQL query will only
// return one record (or none at all, in which case we return a ErrRecordNotFound error).
func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `
		SELECT id, created_at, firstname, lastname, email, password_hash, activated
		FROM users
		WHERE email = $1
	`

	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}
