package service

import rep "real-time-forum/backend/repositories"


type AppService struct {
	repo *rep.AppRepository
}

// NewPostService creates a new service
func NewPostService(repo *rep.AppRepository) *AppService {
	return &AppService{repo: repo}
}
