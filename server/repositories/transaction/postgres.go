package repositories

import (
	"database/sql"
	"go-web-template/server/models"
)

type transactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepo{
		DB: db,
	}
}

func (e *transactionRepo) Save(transaction *models.Transaction) error {
	query := `
		INSERT INTO transactions (
			id, employee_id, menu_id, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5
		)
	`

	stmt, err := e.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		transaction.ID, transaction.Employee.ID, transaction.Menu.ID, transaction.CreatedAt, transaction.UpdatedAt,
	)

	return err
}

func (e *transactionRepo) FindAll() (*[]models.Transaction, error) {
	query := `
		SELECT 
			id, employee_id, menu_id, created_at, updated_at
		FROM
			transactions
	`

	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID, &transaction.Employee.ID, &transaction.Menu.ID, &transaction.CreatedAt, &transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return &transactions, nil
}
