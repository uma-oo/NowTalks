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
	service *service.AppService
}

type UserHanlder struct {
	service *service.AppService
}















func NewCommentHandler(service *service.AppService) *CommentHandler {
	return &CommentHandler{service: service}
}

func NewPostHandler(service *service.AppService) *PostHandler {
	return &PostHandler{service: service}
}

func NewUserHandler(service *service.AppService) *UserHanlder{
	return &UserHanlder{service: service}
}

// NewPostService

func WriteJsonErrors(w http.ResponseWriter, errJson models.ErrorJson) {
	w.WriteHeader(errJson.Status)
	json.NewEncoder(w).Encode(errJson)
}

func WriteDataBack(w http.ResponseWriter, data any) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&data)
}
