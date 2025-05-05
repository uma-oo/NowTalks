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
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *AppService) AddPost(post models.Post) *models.ErrorJson {
	if post.Content == "" || post.Title == "" {
		return &models.ErrorJson{Status: 400, Message: "Bad Request!! Empty title or Message"}
	}
	err := s.repo.CreatePost(&post)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}
