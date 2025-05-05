package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/backend/models"
)

func (handler *AppHandler) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		fmt.Println("error", err)
	}

	

	handler.service.AddPost(post)
}

func (handler *AppHandler) AddPostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		fmt.Println("error", err)
	}

	handler.service.AddPost(post)
}
