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
	fmt.Println("HEREEEE")
	var ErrorJson *models.ErrorJson
	var message *models.Post
	if post.Content == ""  {
		ErrorJson.Status = 400
		message.Content = "ERROR: Empty Body Post!!"
	} else if post.Title ==""{
	  message.Title = "ERROR: Empty Body Title"
	}
	
	if message!= nil {
		ErrorJson.Message = message
		return ErrorJson
	}
	err := s.repo.CreatePost(&post)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}