package main

import (
	"log"

	"github.com/Dmitriihub/tasks-service/internal/database"
	"github.com/Dmitriihub/tasks-service/internal/task"
	transportgrpc "github.com/Dmitriihub/tasks-service/internal/transport/grpc"
)

func main() {
	database.InitDB()
	repo := task.NewRepository(database.DB)
	svc := task.NewService(repo)

	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users service: %v", err)
	}
	defer conn.Close()

	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("failed to run gRPC server: %v", err)
	}
}
