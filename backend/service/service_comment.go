package service

import (
	"fmt"

	"real-time-forum/backend/models"
)




// add the offset and the limit thing after 
func (s *AppService) GetComments(postId int) ([]models.Comment, *models.ErrorJson) {
	comments, err := s.repo.GetComments(postId)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return comments , nil 
}



// check if the content is null 
func (s *AppService) AddComment( comment *models.Comment) *models.ErrorJson {
	if comment.Content=="" {
		return &models.ErrorJson{Status: 500, Message: "Bad request! Comment content is empty!!"}
	}
	if err := s.repo.CreateComment(comment); err!= nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil 
}