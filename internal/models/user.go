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

func (m *UserModel) Insert(user User) error {
	res := m.Db.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (m *UserModel) GetOne(id string) (User, error) {
	var user User
	res := m.Db.First(&user, "Id = ?", id)
	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}
