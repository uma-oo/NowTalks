package handler

import (
	"encoding/json"
	"net/http"
	"real-time-forum/backend/models"
	"real-time-forum/backend/service"
)

type PostHandler struct {
	service *service.AppService
}
 type CommentHandler struct {
	service 
 }





// NewPostService creates a new service
func NewPostService(service *service.AppService) *AppHandler {
	return &AppHandler{service: service}
}

func (ah *AppHandler) AddPost (){}
func (ah *AppHandler) GetPost (){}
func (ah *AppHandler) AddComment (){}
func (ah *AppHandler) GetComment(){}




func WriteJsonErrors(w http.ResponseWriter, errJson models.ErrorJson){
	w.WriteHeader(errJson.Status)
	json.NewEncoder(w).Encode(errJson)
}