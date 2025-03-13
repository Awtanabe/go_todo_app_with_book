package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserID int64

type User struct {
	ID       UserID    `gorm:"column:id;primaryKey"`
	Name     string    `gorm:"column:name"`
	Password string    `gorm:"column:password"`
	Role     string    `gorm:"column:role"`
	Created  time.Time `gorm:"column:created;autoCreateTime"`
	Modified time.Time `gorm:"column:modified;autoUpdateTime"`
}

func (u User) TableName() string {
	return "todo.user"
}

func (u *User) ComparePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}
