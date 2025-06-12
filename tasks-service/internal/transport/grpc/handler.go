package grpc

import (
	"context"
	"fmt"

	taskpb "github.com/Dmitriihub/project-protos/proto/task"
	userpb "github.com/Dmitriihub/project-protos/proto/user"
	"github.com/Dmitriihub/tasks-service/internal/task"
)

type Handler struct {
	svc        *task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *task.Service, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	_, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId})
	if err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}
	t, err := h.svc.CreateTask(task.Task{UserID: req.UserId, Title: req.Title})
	if err != nil {
		return nil, err
	}
	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     t.ID,
			UserId: t.UserID,
			Title:  t.Title,
			IsDone: t.IsDone,
		},
	}, nil
}
