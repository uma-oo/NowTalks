package service

import (
	"time"

	"real-time-forum/backend/models"

	"github.com/google/uuid"
)

// used only for the registration phase
func (s *AppService) SetUserSession(user *models.User) (*models.Session, *models.ErrorJson) {
	session := &models.Session{}
	session.Token = uuid.NewString()
	session.ExpDate = time.Now().Add(24 * time.Hour)
	errJson := s.repo.CreateUserSession(session, user)
	if errJson != nil {
		return nil, errJson
	}
	return session, nil
}

func (s *AppService) GetUserSessionByToken(token string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetUserbyToken(token)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *AppService) CheckUserSession(user *models.User) {
	
}

func (s *AppService) UpdateUserSession(session *models.Session) {
}

func (s *AppService) GetSessionByUserId(user_id int) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetUserSession(user_id)
	if err != nil {
		return nil, err
	}
	return session, nil
}
