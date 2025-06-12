package grpc

import (
	"context"
	"fmt"

	taskpb "github.com/Dmitriihub/project-protos/proto/task"
	userpb "github.com/Dmitriihub/project-protos/proto/user"
	"github.com/Dmitriihub/tasks-service/internal/task"
	"google.golang.org/protobuf/types/known/emptypb"
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
	if _, err := h.userClient.GetUser(ctx, &userpb.User{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}
	t, err := h.svc.CreateTask(task.Task{
		UserID: uint(req.UserId),
		Title:  req.Title,
	})
	if err != nil {
		return nil, err
	}
	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(t.ID),
			UserId: uint32(t.UserID),
			Title:  t.Title,
			IsDone: t.IsDone,
		},
	}, nil
}
func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	t, err := h.svc.GetTaskByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &taskpb.GetTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(t.ID),
			UserId: uint32(t.UserID),
			Title:  t.Title,
			IsDone: t.IsDone,
		},
	}, nil
}

func (h *Handler) ListTasks(ctx context.Context, _ *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, err
	}
	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     uint32(t.ID),
			UserId: uint32(t.UserID),
			Title:  t.Title,
			IsDone: t.IsDone,
		})
	}
	return &taskpb.ListTasksResponse{Tasks: pbTasks}, nil
}

func (h *Handler) ListTasksByUser(ctx context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetTasksByUser(uint(req.UserId))
	if err != nil {
		return nil, err
	}
	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     uint32(t.ID),
			UserId: uint32(t.UserID),
			Title:  t.Title,
			IsDone: t.IsDone,
		})
	}
	return &taskpb.ListTasksResponse{Tasks: pbTasks}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	// Проверка пользователя (опционально, если логика требует)
	if _, err := h.userClient.GetUser(ctx, &userpb.User{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	t, err := h.svc.UpdateTask(task.Task{
		Title:  req.Title,
		IsDone: req.IsDone,
	})
	if err != nil {
		return nil, err
	}
	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(t.ID),
			UserId: uint32(t.UserID),
			Title:  t.Title,
			IsDone: t.IsDone,
		},
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*emptypb.Empty, error) {
	if err := h.svc.DeleteTask(uint(req.Id)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
