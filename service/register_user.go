package service

import (
	"context"
	"fmt"
	"go_todo_app/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUser struct {
	DB   *gorm.DB
	Repo UserRegister
}

func (ru RegisterUser) RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error) {

	// password を []byte(password)でバイナリに変更している
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("cannot hash password %w", &err)
	}

	u := &entity.User{
		Name:     name,
		Password: string(pw),
		Role:     role,
	}

	if err := ru.Repo.RegisterUser(ctx, ru.DB, u); err != nil {
		return nil, fmt.Errorf("failed to register %w", &err)
	}
	return u, nil
}
