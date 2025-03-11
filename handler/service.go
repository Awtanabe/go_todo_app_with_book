package handler

import (
	"context"
	"go_todo_app/entity"
	"go_todo_app/store"
)

type ListTaskService interface {
	ListTask(ctx context.Context, title string) ([]entity.Task, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (entity.Task, error)
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
	tasks, err := ls.taskStore.All()
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (ls *TaskService) AddTask(ctx context.Context, t *entity.Task) (*entity.Task, error) {
	task, err := ls.taskStore.Add(t)
	if err != nil {
		return nil, err
	}

	return task, nil
}
