package main

import (
	"log"

	"github.com/Dmitriihub/tasks-service/internal/database"
	"github.com/Dmitriihub/tasks-service/internal/task"
	transportgrpc "github.com/Dmitriihub/tasks-service/internal/transport/grpc"
)

func main() {
	// 1. Инициализация базы данных
	database.InitDB()

	// 2. Репозиторий и сервис задач
	repo := task.NewRepository(database.DB)
	svc := task.NewService(repo)

	// 3. gRPC-клиент Users-сервиса
	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Запуск gRPC-сервера Tasks-сервиса
	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
