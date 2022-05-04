package repositories

import "go-web-template/server/models"

type TransactionRepository interface {
	Save(transaction *models.Transaction) error
	FindAll() (*[]models.Transaction, error)
}
