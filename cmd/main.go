package main

import (
	"log"
	"os"

	"github.com/MigueelLago/todo-go/internal/config"
	"github.com/MigueelLago/todo-go/internal/interfaces/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	db := config.InitDB()
	router := http.SetupRoutes(db)
	router.Run(":" + port)
	log.Println("Servidor rodando na porta " + port + "...")
}
