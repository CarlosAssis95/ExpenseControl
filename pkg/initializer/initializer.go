package initializer

import (
	config2 "Controle-Despesas/internal/config"
	database "Controle-Despesas/internal/infrastructure"
	"log"
)

func Boo() {

	config, err := config2.Load()
	if err != nil {
		log.Println("Error in start config. ", err)
	}

	db, err := database.ConnectDatabase(config.Database)
	if err != nil {
		log.Println("error in start db. ", err)
	}

	defer db.Close()

}
