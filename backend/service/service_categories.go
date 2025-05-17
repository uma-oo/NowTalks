package service

import "real-time-forum/backend/models"

func (s *AppService) GetAllCategories() ([]models.Category, *models.ErrorJson) {
	categories, err := s.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}


