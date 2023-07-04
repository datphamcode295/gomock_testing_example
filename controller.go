package controller

import (
	"errors"

	"example.com/testing/models"
	"example.com/testing/store"
	"gorm.io/gorm"
)

type Auth struct {
	UserRepository store.UserRepository
}

func (c *Auth) Login(email string, pass string) (*models.User, error) {
	user, err := c.UserRepository.GetUser(email, pass)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("invalid email or password")
		}

		return nil, err
	}

	return user, nil
}
