package models

import (
	"time"
)

type Notification struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"`
	Data      string    `json:"data"`
	ReadAt    time.Time `json:"read_at,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
