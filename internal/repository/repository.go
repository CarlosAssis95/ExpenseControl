package repository

import (
	"database/sql"
	_ "embed"
)

var (
	//go:embed querys/queryInsertExpense.sql
	queryInsertExpense string
)

type RepositoryExpense struct {
	db *sql.DB
}

type RepositoryExpenseInterface interface {
}

func NewRepositoryExpense(db *sql.DB) *RepositoryExpense {
	return &RepositoryExpense{db: db}
}

func (r *RepositoryExpense) InsertExpense() error {

	_, err := r.db.Exec(queryInsertExpense)
	if err != nil {
		return err
	}

	return nil
}
