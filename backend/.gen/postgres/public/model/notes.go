//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Notes struct {
	ID         uuid.UUID `sql:"primary_key"`
	UserID     uuid.UUID
	Title      string
	Content    *string
	Icon       *string
	IsArchived bool
	ParentID   *uuid.UUID
	CoverImage *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
