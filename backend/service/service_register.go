package service

import (
	"real-time-forum/backend/models"
)

func (s *AppService) Register(user *models.User) (*models.User, error) {
	return user, nil
}

// Section of the helper functions used to check the user input

// check if the NICKNAME is duplicated

