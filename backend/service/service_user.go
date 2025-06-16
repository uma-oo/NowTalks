package service

import (
	"net/http"

	"real-time-forum/backend/models"
)

func (s *AppService) GetUser(login *models.Login) (*models.User, *models.ErrorJson) {
	user, err := s.repo.GetUser(login)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Message: err.Message}
	}
	return user, nil
}

func (s *AppService) GetUsers(offset, user_id int) ([]models.User, *models.ErrorJson) {
	users, err := s.repo.GetUsers(offset, user_id)
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

func (service *AppService) GetUsernameFromSession(r *http.Request) int {
	cookie, _ := r.Cookie("session")
	session, _ := service.GetSessionByTokenEnsureAuth(cookie.Value)
	return session.UserId
}
