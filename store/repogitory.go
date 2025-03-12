package store

import (
	"context"
	"fmt"
	"go_todo_app/config"

	"github.com/budougumi0617/go_todo_app/clock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	Clocker clock.Clocker
}

func New(ctx context.Context, cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
