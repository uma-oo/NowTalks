package service

import m "real-time-forum/backend/models"

func (s *AppService) IsLoggedInUser(token string) (*m.UserData, *m.ErrorJson) {
	user, err := s.repo.IsLoggedInUser(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}
