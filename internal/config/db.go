package config

import (
	"log"

	"github.com/MigueelLago/todo-go/internal/domain/task"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	err = db.AutoMigrate(&task.Task{})
	if err != nil {
		log.Fatal("Erro ao migrar o banco de dados")
	}

	return db
}
