package note

import (
	"time"

	"backend/api"

	"github.com/google/uuid"
)

type Note struct {
	Id         *uuid.UUID      `json:"id"`
	UserId     *uuid.UUID      `json:"user_id"`
	Title      string          `json:"title"`
	Content    *api.Content    `json:"content"`
	IsArchived bool            `json:"is_archived"`
	ParentId   *uuid.UUID      `json:"parent_id"`
	CoverImage *api.CoverImage `json:"cover_image"`
	Icon       *api.Icon       `json:"icon"`
	CreatedAt  *time.Time      `json:"created_at"`
	UpdatedAt  *time.Time      `json:"updated_at"`
}
