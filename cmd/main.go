package main

import (
	"log"

	"github.com/MigueelLago/todo-go/internal/config"
	"github.com/MigueelLago/todo-go/internal/interfaces/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	db := config.InitDB()
	router := http.SetupRoutes(db)
	router.Run(":3000")
	log.Println("Servidor rodando na porta 3000")
}
