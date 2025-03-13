package service

import (
	"context"
	"fmt"
	"go_todo_app/handler"

	"gorm.io/gorm"
)

type Login struct {
	DB             *gorm.DB
	Repo           handler.UserGetter
	TokenGenerator handler.TokenGenerator
}

func (l *Login) Login(ctx context.Context, name, pw string) (string, error) {
	// nameだけだけどしたでpassも検証
	u, err := l.Repo.GetUser(ctx, l.DB, name)

	if err != nil {
		return "", fmt.Errorf("failed to list %w", err)
	}

	if err = u.ComparePassword(pw); err != nil {
		return "", fmt.Errorf("wrong passwird %w", err)
	}

	jwt, err := l.TokenGenerator.GenerateToken(ctx, *u)

	if err != nil {
		return "", fmt.Errorf("field to generate JWT %w", err)
	}

	return string(jwt), nil
}
