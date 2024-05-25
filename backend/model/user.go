package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
