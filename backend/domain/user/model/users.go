package model

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	ID        uuid.UUID `sql:"primary_key"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
