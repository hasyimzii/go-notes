package models

import (
	"time"
)

type NoteUser struct {
	ID        uint      `json:"id"`
	NoteID    uint      `json:"note_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
