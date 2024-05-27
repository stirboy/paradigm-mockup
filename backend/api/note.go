package api

import "github.com/google/uuid"

type Note struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"userId"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	IsArchived bool      `json:"isArchived"`
	ParentId   uuid.UUID `json:"parentId"`
	CoverImage string    `json:"coverImage"`
	Icon       string    `json:"icon"`
}

type Content string
type CoverImage string
type Icon string
