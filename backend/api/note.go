package api

import (
	"github.com/google/uuid"
	"github.com/oapi-codegen/nullable"
)

type Note struct {
	Id         uuid.UUID  `json:"id"`
	UserId     uuid.UUID  `json:"userId"`
	Title      string     `json:"title"`
	Content    Content    `json:"content,omitempty"`
	IsArchived bool       `json:"isArchived"`
	ParentId   *uuid.UUID `json:"parentId,omitempty"`
	CoverImage CoverImage `json:"coverImage,omitempty"`
	Icon       Icon       `json:"icon,omitempty"`
}

type Content struct {
	nullable.Nullable[string]
}

func (c Content) MarshalJSON() ([]byte, error) {
	return c.Nullable.MarshalJSON()
}

type CoverImage struct {
	nullable.Nullable[string]
}

func (c CoverImage) MarshalJSON() ([]byte, error) {
	return c.Nullable.MarshalJSON()
}

type Icon struct {
	nullable.Nullable[string]
}

func (icon Icon) MarshalJSON() ([]byte, error) {
	return icon.Nullable.MarshalJSON()
}
