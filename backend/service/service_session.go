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

func (s *AppService) GetSessionByTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetSessionbyTokenEnsureAuth(token)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *AppService) CheckUserSession(token string) (bool, *models.Session) {
	has_session, session := s.repo.HasValidToken(token)
	if has_session {
		return true, session
	}
	return false, nil
}

func (s *AppService) UpdateUserSession(session *models.Session) (*models.Session, *models.ErrorJson) {
	new_session := models.NewSession()
	new_session.Token = uuid.NewString()
	new_session.ExpDate = time.Now().Add(24 * time.Hour)
	if err := s.repo.UpdateSession(session, new_session); err != nil {
		return nil, err
	}
	return new_session, nil
}

func (s *AppService) GetSessionByUserId(user_id int) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetUserSessionByUserId(user_id)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *AppService) DeleteSession(session *models.Session) *models.ErrorJson {
	if err := s.repo.DeleteSession(*session); err != nil {
		return err
	}
	return nil
}
