package services

import (
	"database/sql"
	"go-web-template/server/helper"
	"go-web-template/server/models"
	params "go-web-template/server/params/transaction"
	repositories "go-web-template/server/repositories/transaction"
	"time"
)

type TransactionServices struct {
	TransactionRepository repositories.TransactionRepository
	DB                    *sql.DB
}

func NewTransactionService(db *sql.DB) *TransactionServices {
	repositories := repositories.NewTransactionRepository(db)
	return &TransactionServices{
		TransactionRepository: repositories,
		DB:                    db,
	}
}

func (t *TransactionServices) CreateNewTransaction(request *params.TransactionCreate) bool {
	defer helper.HandleError()

	emp := request.ParseToModel()
	err := t.TransactionRepository.Save(emp)

	if err != nil {
		helper.HandlePanicIfError(err)
		return false
	}

	return true
}

func (t *TransactionServices) GetAllTransactions() *[]params.TransactionSingleView {
	defer helper.HandleError()
	transactions, err := t.TransactionRepository.FindAll()
	helper.HandlePanicIfError(err)
	return makeTransactionListView(transactions)
}

func makeTransactionListView(models *[]models.Transaction) *[]params.TransactionSingleView {
	var transactionListView []params.TransactionSingleView
	for _, model := range *models {
		transactionListView = append(transactionListView, *makeTransactionSingleView(&model))
	}
	return &transactionListView
}

func makeTransactionSingleView(models *models.Transaction) *params.TransactionSingleView {
	return &params.TransactionSingleView{
		ID:        models.ID,
		Employee:  models.Employee.Name,
		Menu:      models.Menu.Name,
		CreatedAt: models.CreatedAt.Format(time.RFC3339),
		UpdatedAt: models.UpdatedAt.Format(time.RFC3339),
	}
}
