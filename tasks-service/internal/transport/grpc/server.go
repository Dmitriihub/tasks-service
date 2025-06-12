package grpc

import (
	"net"

	taskpb "github.com/Dmitriihub/project-protos/proto/task"
	"github.com/Dmitriihub/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc *task.Service, uc taskpb.UserServiceClient) error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	handler := NewHandler(svc, uc)
	taskpb.RegisterTaskServiceServer(grpcServer, handler)

	return grpcServer.Serve(lis)
}
