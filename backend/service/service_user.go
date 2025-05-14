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






func (s *AppService) IsLoggedIn() bool {
	return true
}
