package data

import (
	"time"
)

type Movie struct {
	ID         int64 `json:"id"`
	CreatedAt  time.Time
	LastEdited time.Time
	Content    string
}
