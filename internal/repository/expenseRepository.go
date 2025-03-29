package repository

import (
	"Controle-Despesas/internal/domain"
	"database/sql"
	_ "embed"
)

var (
	//go:embed query/queryInsertExpense.sql
	queryInsertExpense string
)

type ExpenseRepository struct {
	db *sql.DB
}

type ExpenseRepositoryInterface interface {
	InsertExpense(expenseDomain domain.Expense) error
}

func NewRepositoryExpense(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) InsertExpense(expenseDomain domain.Expense) error {

	_, err := r.db.Exec(
		queryInsertExpense,
		expenseDomain.DescriptionExpense,
		expenseDomain.QuantityInstallments,
		expenseDomain.ValueInstallments,
		expenseDomain.DatePaymentExpense)
	if err != nil {
		return err
	}

	return nil
}
