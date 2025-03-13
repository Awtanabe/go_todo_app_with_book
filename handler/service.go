package handler

import (
	"context"
	"fmt"
	"go_todo_app/auth"
	"go_todo_app/entity"
	"go_todo_app/store"
	"time"

	"gorm.io/gorm"
)

type ListTaskService interface {
	ListTask(ctx context.Context, title string) ([]entity.Task, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}

type LoginService interface {
	Login(ctx context.Context, name, pw string) (string, error)
}

type UserGetter interface {
	GetUser(ctx context.Context, db *gorm.DB, name string) (*entity.User, error)
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u entity.User) ([]byte, error)
}

type TaskService struct {
	taskStore *store.TaskStore
}

func NewTaskService(taskStore *store.TaskStore) *TaskService {
	return &TaskService{taskStore}
}

func NewAddService(taskStore *store.TaskStore) *TaskService {
	return &TaskService{taskStore}
}

func (ls *TaskService) ListTask(ctx context.Context) ([]entity.Task, error) {
	id, ok := auth.GetUserID(ctx)

	if !ok {
		return nil, fmt.Errorf("user id not found")
	}
	tasks, err := ls.taskStore.All(id)
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (ls *TaskService) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	t := &entity.Task{
		Title:    title,
		UserID:   id,
		Status:   entity.TaskStatusTodo,
		Created:  time.Now(),
		Modified: time.Now(),
	}

	task, err := ls.taskStore.Add(t)
	if err != nil {
		return nil, err
	}

	return task, nil
}
