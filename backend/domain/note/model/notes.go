package model

import (
	"github.com/google/uuid"
	"time"
)

// Notes struct represents notes table in database
type Notes struct {
	ID         uuid.UUID  `sql:"primary_key" alias:"notes.id"`
	UserID     uuid.UUID  `alias:"notes.user_id"`
	Title      string     `alias:"notes.title"`
	Content    *string    `alias:"notes.content"`
	Icon       *string    `alias:"notes.icon"`
	IsArchived bool       `alias:"notes.is_archived"`
	ParentID   *uuid.UUID `alias:"notes.parent_id"`
	CoverImage *string    `alias:"notes.cover_image"`
	CreatedAt  time.Time  `alias:"notes.created_at"`
	UpdatedAt  time.Time  `alias:"notes.updated_at"`
}
