package params

import (
	"go-web-template/server/models"
)

type TransactionCreate struct {
	EmployeeID string `json:"employee_id"`
	MenuID     string `json:"menu_id"`
}

func (t *TransactionCreate) ParseToModel() *models.Transaction {
	transaction := models.NewTransaction()
	transaction.Employee.Name = t.EmployeeID
	transaction.Menu.Name = t.MenuID
	return transaction
}
