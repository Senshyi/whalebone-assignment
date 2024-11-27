package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Db *gorm.DB
}

type User struct {
	ID          string `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Email       string
	DateOfBirth time.Time
}

func (m *UserModel) Insert() error {
	return nil
}
func (m *UserModel) GetOne(id string) (User, error) {
	return User{}, nil
}
