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

func (ah *AppHandler) AddPost (){}
func (ah *AppHandler) GetPost (){}
func (ah *AppHandler) AddComment (){}
func (ah *AppHandler) GetComment(){}
