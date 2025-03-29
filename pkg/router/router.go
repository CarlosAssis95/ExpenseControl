package router

import (
	"Controle-Despesas/internal/controller"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router) {
	r.HandleFunc("/receive", controller.InsertExpense).Methods("POST")
}
