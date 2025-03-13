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

func (r *Repository) GetUser(
	ctx context.Context, db *gorm.DB, name string,
) (*entity.User, error) {
	u := &entity.User{}
	if err := db.WithContext(ctx).Where("name = ?", name).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
