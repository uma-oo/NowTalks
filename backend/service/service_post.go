package service

import (
	"fmt"
	"strings"

	"real-time-forum/backend/models"
	"real-time-forum/backend/utils"
)

// ayoub u afkaru lghariba
// bash ghan3mr hadshi :)
// add offsets and limits

func (s *AppService) AddPost(post *models.Post) (*models.Post, *models.ErrorJson) {
	errorJson := models.NewErrorJson(0, "")
	message := models.NewPostErr()
	if strings.TrimSpace(post.Content) == "" {
		message.Content = "empty post content!!"
	}
	if  strings.TrimSpace(post.Title) == "" {
		message.Title = "empty title content!!"
	}
	if len(post.PostCategories) == 0 || !utils.CheckPOSTCategories(post.PostCategories) {
		message.Categories = "please choose at least one category!"
	}
	if message.Content != "" || message.Title != "" || message.Categories != "" {
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

func (s *AppService) GetPosts(user_id, offset int) ([]models.Post, *models.ErrorJson) {
	posts, err := s.repo.GetPosts(user_id, offset)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
