package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"real-time-forum/backend/models"
)

// OMG

func (appRep *AppRepository) CreatePost(post *models.Post) (*models.Post, *models.ErrorJson) {
	post_created := models.NewPost()
	query := `INSERT INTO posts(userID,  title, content) VALUES (?, ?, ?) RETURNING postID, title , content ,createdAt, total_comments`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.UserId, post.Title, post.Content).Scan(&post_created.Id, &post_created.Title,
		&post_created.Content, &post_created.CreatedAt, &post_created.TotalComments)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	username, errJson := appRep.getUserNameById(post.UserId)
	if errJson != nil {
		return nil, errJson
	}
	post_created.Username = username
	return post_created, nil
}

// all the posts
// add the offset and the limit after
func (appRep *AppRepository) GetPosts(limit, offset int) ([]models.Post, *models.ErrorJson) {
	var posts []models.Post
	query := `SELECT  users.nickname, posts.createdAt, posts.title, posts.content FROM posts 
	INNER JOIN users 
	ON posts.userID = users.userID
	ORDER BY posts.createdAt DESC
	LIMIT ?
	OFFSET ?;
	`
	rows, err := appRep.db.Query(query, limit, offset)
	if rows.Err() == sql.ErrNoRows {
		return posts, nil
	}
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Username, &post.CreatedAt, &post.Title, &post.Content); err != nil {
			return posts, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		posts = append(posts, post)

	}
	defer rows.Close()
	return posts, nil
}

// Filter by MyPosts // by userId
// Filter based on categories
// based on the len of the category
func (appRep *AppRepository) GetPostsByCategory(limit, offset int, category ...string) ([]models.Post, *models.ErrorJson) {
	posts := []models.Post{}
	if len(category) == 0 {
		return appRep.GetPosts(limit, offset)
	}
	categories := strings.Join(category, ",")
	fmt.Println("categories", categories)
	query := fmt.Sprintf(`SELECT * 
	FROM posts INNER JOIN postCategories 
	ON posts.postID = postCategories.postID
	INNER JOIN categories ON postCategories.categoryID = categories.categoryID
	WHERE categories.category IN (%v)
	ORDER BY posts.createdAt DESC
	LIMIT ? OFFSET ?`,
		categories)
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	rows, err := stmt.Query(query, limit, offset)
	if rows.Err() == sql.ErrNoRows {
		return posts, nil
	}
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post); err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		posts = append(posts, post)
	}
	defer rows.Close()

	return posts, nil
}

func (appRep *AppRepository) GetPostsByUser(user_id, offset, limit int) ([]models.Post, *models.ErrorJson) {
	var posts []models.Post
	query := `SELECT  users.nickname, posts.createdAt, posts.title, posts.content FROM posts 
	INNER JOIN users 
	ON posts.userID = users.userID
	WHERE users.userID = ?
	ORDER BY posts.createdAt DESC
	LIMIT ?
	OFFSET ?;
	`
	rows, err := appRep.db.Query(query, user_id, limit, offset)
	if rows.Err() == sql.ErrNoRows {
		return posts, nil
	}
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Username, &post.CreatedAt, &post.Title, &post.Content); err != nil {
			return posts, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		posts = append(posts, post)

	}
	defer rows.Close()
	return posts, nil
}
