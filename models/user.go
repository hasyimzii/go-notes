package models

import (
	"time"
)

type User struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Phone      string    `json:"phone,omitempty"`
	Meta       string    `json:"meta,omitempty"`
	VerifiedAt time.Time `json:"verified_at,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
