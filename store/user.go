package store

import (
	"context"
	"go_todo_app/entity"

	"gorm.io/gorm"
)

func (r *Repository) RegisterUser(ctx context.Context, db *gorm.DB, u *entity.User) error {

	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil
}
