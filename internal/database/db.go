package database

import (
	"log"

	"github.com/Dmitriihub/tasks-service/internal/task"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	if err := DB.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("ошибка миграции: %v", err)
	}
}
