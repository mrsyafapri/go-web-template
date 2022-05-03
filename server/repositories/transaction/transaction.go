package repositories

import "go-web-template/server/models"

type TransactionRepository interface {
	Save(transaction *models.Transaction) error
	FindAll() (*[]models.Transaction, error)
	FindByID(id string) (*models.Transaction, error)
	UpdateByID(transaction *models.Transaction) error
	DeleteByID(id string) error
}
