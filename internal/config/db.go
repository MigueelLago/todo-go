package config

import (
	"log"
	"os"

	"github.com/MigueelLago/todo-go/internal/domain/task"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "tasks.db"
	}
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	err = db.AutoMigrate(&task.Task{})
	if err != nil {
		log.Fatal("Erro ao migrar o banco de dados")
	}

	return db
}
