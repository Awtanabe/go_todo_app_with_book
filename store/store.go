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
	db *gorm.DB
}

func NewTaskStore(db *gorm.DB) *TaskStore {
	return &TaskStore{db: db}
}

func (ts *TaskStore) Add(t *entity.Task) (*entity.Task, error) {
	if err := ts.db.Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (ts *TaskStore) All(id entity.UserID) (entity.Tasks, error) {
	tasks := []entity.Task{}

	if err := ts.db.Where("user_id = ?", id).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
