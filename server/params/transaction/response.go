package params

import (
	"github.com/google/uuid"
)

type TransactionSingleView struct {
	ID        uuid.UUID
	Employee  string
	Menu      string
	CreatedAt string
	UpdatedAt string
}
