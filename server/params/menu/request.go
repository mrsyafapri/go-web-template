package params

import (
	"go-web-template/server/models"
	"time"

	"github.com/google/uuid"
)

type MenuCreate struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
}

func (m *MenuCreate) ParseToModel() *models.Menu {
	menu := models.NewMenu()
	menu.Category = m.Category
	menu.Name = m.Name
	menu.Desc = m.Desc
	return menu
}

type MenuUpdate struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Desc      string    `json:"desc"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *MenuUpdate) ParseToModel() *models.Menu {
	return &models.Menu{
		ID:        m.ID,
		Name:      m.Name,
		Category:  m.Category,
		Desc:      m.Desc,
		UpdatedAt: time.Now(),
	}
}
