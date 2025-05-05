package repositories

import (
	"fmt"

	"real-time-forum/backend/models"
)

// OMG

func (appRep *AppRepository) CreatePost(post *models.Post) *models.ErrorJson {
	fmt.Println("inside repo")
	query := `INSERT INTO posts(userID,  title, content) VALUES (?, ?, ?)`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	_, err = stmt.Exec(post.UserId, post.Title, post.Content)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// all the posts
func (appRep *AppRepository) GetPosts() ([]models.Post, *models.ErrorJson) {
	var posts []models.Post
	query := `SELECT postID , userID, createdAt, title, content FROM posts`
	rows, err := appRep.db.Query(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.CreatedAt, &post.Title, &post.Content); err != nil {
			return posts,  &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		posts = append(posts, post)

	}
	return posts, nil
}

// Filter by MyPosts // by userId
// Filter based on categories
