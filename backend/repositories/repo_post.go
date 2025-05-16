package repositories

import (
	"fmt"

	"real-time-forum/backend/models"
)

// OMG

func (appRep *AppRepository) CreatePost(post *models.Post) (*models.Post, *models.ErrorJson) {
	post_created := models.NewPost()
	query := `INSERT INTO posts(userID,  title, content) VALUES (?, ?, ?) RETURNING *`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.UserId, post.Title, post.Content).Scan(&post_created.Id, &post_created.UserId,
		&post_created.CreatedAt, &post_created.Title,
		&post_created.Content, &post_created.TotalComments)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	username, errJson := appRep.getUserNameById(post_created.UserId)
	if errJson != nil {
		return nil, errJson
	}
	post_created.Username = username
	return post_created, nil
}

// all the posts
// add the offset and the limit after
func (appRep *AppRepository) GetPosts(limit int, offset int) ([]models.Post, *models.ErrorJson) {
	var posts []models.Post
	query := `SELECT posts.postID , users.nickname, posts.createdAt, posts.title, posts.content FROM posts 
	INNER JOIN users 
	ON posts.userID = users.userID
	ORDER BY posts.createdAt DESC
	LIMIT ?
	OFFSET ?;
	`
	rows, err := appRep.db.Query(query, limit, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Username, &post.CreatedAt, &post.Title, &post.Content); err != nil {
			return posts, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		posts = append(posts, post)

	}
	defer rows.Close()
	return posts, nil
}

// Filter by MyPosts // by userId
// Filter based on categories

// func (appRep  *AppRepository) GetPostsByCategory(category... string) ([]models.Post , *models.ErrorJson){
// 	var posts = []models.Post{}
// 	query := Query{}
// 	query.Query("SELECT * FROM postCategories ")

// }
