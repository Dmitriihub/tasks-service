package grpc

import (
	"net"

	taskpb "github.com/Dmitriihub/project-protos/proto/task"
	userpb "github.com/Dmitriihub/project-protos/proto/user"
	"github.com/Dmitriihub/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc *task.Service, uc userpb.UserServiceClient) error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}
	grpcSrv := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(grpcSrv, NewHandler(svc, uc))
	return grpcSrv.Serve(lis)
}
