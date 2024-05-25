package model

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id        uuid.UUID  `json:"id"`
	UserId    uuid.UUID  `json:"user_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
