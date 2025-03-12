package service

import (
	"context"
	"go_todo_app/entity"

	"gorm.io/gorm"
)

type UserRegister interface {
	RegisterUser(ctx context.Context, db *gorm.DB, u *entity.User) error
}
