package repositories

import "real-time-forum/backend/models"



// OMG

func (appRep *AppRepository) CreatePost(post *models.Post) error {
	query := `INSERT INTO posts VALUES (?, ?, ?, ?)`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.UserId, post.CreatedAt, post.Title, post.Content)
	if err != nil {
		return err
	}
	return nil
}



// all the posts 
func (appRep *AppRepository) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT postID , userID, createdAt, title, content FROM posts`
	rows, err := appRep.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.CreatedAt, &post.Title, &post.Content); err != nil {
			return posts, err
		}
		posts = append(posts, post)

	}
	return posts, nil
}

// Filter by MyPosts // by userId 
// Filter based on categories 

