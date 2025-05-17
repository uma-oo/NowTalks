package service

import (
	"fmt"

	"real-time-forum/backend/models"
)

// ayoub u afkaru lghariba
// bash ghan3mr hadshi :)
// add offsets and limits

func (s *AppService) GetPosts(limit int, offset int) ([]models.Post, *models.ErrorJson) {
	posts, err := s.repo.GetPosts(limit, offset)
	fmt.Println("posts inside the service", posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *AppService) AddPost(post *models.Post) (*models.Post, *models.ErrorJson) {
	errorJson := models.NewErrorJson(0, "")
	message := models.NewPostErr()
	if post.Content == "" {
		message.Content = "ERROR: Empty Post Content!!"
	}
	if post.Title == "" {
		message.Title = "ERROR: Empty Title Content!!"
	}
	if message.Content != "" || message.Title != "" {
		errorJson.Status = 400
		errorJson.Message = message
		return nil, errorJson
	}
	post_created, err := s.repo.CreatePost(post)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return post_created, nil
}

func (s *AppService) GetPostsByCategory(limit int, offset int, category ...string) ([]models.Post, *models.ErrorJson) {
	posts, errJson := s.repo.GetPostsByCategory(limit, offset, category...)
	if errJson != nil {
		return nil, errJson
	}
	return posts, nil
}
