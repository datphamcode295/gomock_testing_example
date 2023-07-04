package store

import (
	"example.com/testing/models"
)

type UserRepository interface {
	GetUser(email string, pass string) (*models.User, error)
	UpdateUser(user *models.User) error
}
