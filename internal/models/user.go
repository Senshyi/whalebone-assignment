package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Db *gorm.DB
}

type DatabaseUser struct {
	ID          string `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Email       string
	DateOfBirth time.Time
}

type ResponseUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}

func DatabaseUserToResponseUser(user DatabaseUser) ResponseUser {
	return ResponseUser{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Name,
		DateOfBirth: user.DateOfBirth.Format("2006-01-02T15:04:05-07:00"),
	}
}

func (m *UserModel) Insert(user DatabaseUser) error {
	res := m.Db.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (m *UserModel) GetOne(id string) (DatabaseUser, error) {
	var user DatabaseUser
	res := m.Db.First(&user, "Id = ?", id)
	if res.Error != nil {
		return DatabaseUser{}, res.Error
	}
	return user, nil
}
