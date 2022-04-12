package params

import (
	"github.com/google/uuid"
)

type EmployeeSingleView struct {
	ID        uuid.UUID
	NIP       string
	Name      string
	Address   string
	CreatedAt string
	UpdatedAt string
}
