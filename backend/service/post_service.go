package service

import (
	"fmt"

	"real-time-forum/backend/models"
)

// ayoub u afkaru lghariba
// bash ghan3mr hadshi :)

func (s *AppService) GetPosts() ([]models.Post, error) {
	posts, err := s.repo.GetPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *AppService) AddPost(post models.Post) error {
	fmt.Println("post", post.Content, post.Title, post.UserId)
	err := s.repo.CreatePost(&post)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}
	return nil
}
