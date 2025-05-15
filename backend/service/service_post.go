package service

import (
	"fmt"

	"real-time-forum/backend/models"
)

// ayoub u afkaru lghariba
// bash ghan3mr hadshi :)
// add offsets and limits

func (s *AppService) GetPosts() ([]models.Post, *models.ErrorJson) {
	posts, err := s.repo.GetPosts()
    fmt.Println("posts inside the service", posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *AppService) AddPost(post *models.Post) *models.ErrorJson {
	ErrorJson := models.NewErrorJson(0, "")
	message := models.NewPostErr()
	if post.Content == "" {
		message.Content = "ERROR: Empty Post Content!!"
	}
	if post.Title == "" {
		message.Title = "ERROR: Empty Title Content!!"
	}
	if message.Content != "" || message.Title != "" {
		ErrorJson.Status = 400
		ErrorJson.Message = message
		return ErrorJson
	}
	err := s.repo.CreatePost(post)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}
