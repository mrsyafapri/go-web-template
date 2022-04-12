package repositories

import "go-web-template/server/models"

type EmployeeRepository interface {
	Save(employee *models.Employee) error
	FindAll() (*[]models.Employee, error)
	FindByID(id string) (*models.Employee, error)
	UpdateByID(employee *models.Employee) error
	DeleteByID(id string) error
}
