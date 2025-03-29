package initializer

import (
	config2 "Controle-Despesas/internal/config"
	"Controle-Despesas/internal/controller"
	database "Controle-Despesas/internal/infrastructure"
	"Controle-Despesas/internal/repository"
	"Controle-Despesas/pkg/router"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Boot() {

	config, err := config2.Load()
	if err != nil {
		log.Println("Error in start config. ", err)
	}

	r := mux.NewRouter()
	router.NewRouter(r)

	db, err := database.ConnectDatabase(config.Database)
	if err != nil {
		log.Println("error in start db. ", err)
	}

	defer db.Close()

	repo := repository.NewRepositoryExpense(db.DB())
	controller.SetExpenseRepository(repo)

	port := os.Getenv("PORT")

	if err := http.ListenAndServe(port, r); err != nil {
		log.Println("Error stating server")
	}
}
