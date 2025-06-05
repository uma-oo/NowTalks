package service

import (
	"fmt"

	"real-time-forum/backend/models"
)

func (s *AppService) GetUser(login *models.Login) (*models.User, *models.ErrorJson) {
	user, err := s.repo.GetUser(login)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return user, nil
}

func (s *AppService) GetUsers(offset int) ([]models.User, *models.ErrorJson) {
	users, err := s.repo.GetUsers(offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *AppService) UserExists(id int) (bool, *models.ErrorJson) {
	exists, errJson := s.repo.UserExists(id)
	if errJson != nil {
		return false, errJson
	}
	return exists, nil
}
