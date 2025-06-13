package service

import (
	"fmt"
	"strings"

	"real-time-forum/backend/models"
)

// add the offset and the limit thing after
func (s *AppService) GetComments(user_id,postId, offset int) ([]models.Comment, *models.ErrorJson) {
	comments, err := s.repo.GetComments(user_id,postId, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return comments, nil
}

// check if the content is null
func (s *AppService) AddComment(comment *models.Comment) (*models.Comment, *models.ErrorJson) {
	message := models.NewCommentErr()
	if strings.TrimSpace(comment.Content) == "" {
		message.Content = "ERROR: Empty Body Comment!"
	}
	if comment.PostId == 0 {
		message.PostId = "ERROR: Post ID is incorrect or did you mean post_id?"
	}
	if message.Content != "" || message.PostId != "" {
		return nil, &models.ErrorJson{Status: 400, Message: message}
	}
	comment_created, err := s.repo.CreateComment(comment)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return comment_created, nil
}
