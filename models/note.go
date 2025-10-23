package models

import (
	"time"
)

type Note struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	Image        string    `json:"image,omitempty"`
	ReminderTime time.Time `json:"reminder_time,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ImageURL     string    `json:"image_url,omitempty"`
	Users        []User    `gorm:"many2many:note_users;"`
}
