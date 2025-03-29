package controller

import (
	"Controle-Despesas/internal/domain"
	"Controle-Despesas/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

var (
	expenseRepository repository.ExpenseRepositoryInterface
)

func SetExpenseRepository(e repository.ExpenseRepositoryInterface) {
	expenseRepository = e
}

func InsertExpense(w http.ResponseWriter, r *http.Request) {

	var expense domain.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		log.Println("Error receiving request body")
	}

	err = expenseRepository.InsertExpense(expense)
	if err != nil {
		log.Println("Error in insert on database")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"message": "Expense recorded successfully",
		"expense": expense,
	}

	if err != nil {
		log.Println("Expense recorded successfully")
		json.NewEncoder(w).Encode(response)
	}
}
