package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID        uuid.UUID
	NIP       string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEmployee() *Employee {
	return &Employee{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
