package store

import (
	"errors"
	"go_todo_app/entity"
)

var (
	Tasks       = &TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}
	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make(entity.Tasks, 0, len(ts.Tasks)) // スライスの正しい初期化

	for _, t := range ts.Tasks { // i ではなく、順番関係なく append する
		tasks = append(tasks, t)
	}

	return tasks
}
