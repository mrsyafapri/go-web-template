package params

import (
	"github.com/google/uuid"
)

type MenuSingleView struct {
	ID        uuid.UUID
	Name      string
	Category  string
	Desc      string
	CreatedAt string
	UpdatedAt string
}
