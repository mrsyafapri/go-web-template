package repositories

import "go-web-template/server/models"

type MenuRepository interface {
	Save(menu *models.Menu) error
	FindAll() (*[]models.Menu, error)
	FindByID(id string) (*models.Menu, error)
	UpdateByID(menu *models.Menu) error
	DeleteByID(id string) error
}
