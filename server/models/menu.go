package models

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID        uuid.UUID
	Name      string
	Category  string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMenu() *Menu {
	return &Menu{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
