package store

import (
	"context"
	"fmt"
	"go_todo_app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(ctx context.Context, cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
