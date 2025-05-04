package handler

import (
	"real-time-forum/backend/service"
)

type AppHandler struct {
	service *service.AppService
}

// NewPostService creates a new service
func NewPostService(service *service.AppService) *AppHandler {
	return &AppHandler{service: service}
}
