package store

import (
	"errors"
	"go_todo_app/entity"

	"gorm.io/gorm"
)

var (
	Tasks       = &TaskStore{db: &gorm.DB{}}
	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskID
	db     *gorm.DB
}

func NewTaskStore(db *gorm.DB) *TaskStore {
	return &TaskStore{db: db}
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	return t.ID, nil
}

func (ts *TaskStore) All() (entity.Tasks, error) {
	tasks := []entity.Task{}

	if err := ts.db.Find(&tasks).Error; err != nil {

		return nil, err
	}

	return tasks, nil
}
