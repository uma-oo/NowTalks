package repositories

import (
	"database/sql"
	"fmt"

	"real-time-forum/backend/models"
)

// OMG

func (appRep *AppRepository) CreatePost(post *models.Post) (*models.Post, *models.ErrorJson) {
	post_created := models.NewPost()
	query := `INSERT INTO posts(userID,  title, content) VALUES (?, ?, ?) RETURNING postID, title , content ,createdAt`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.UserId, post.Title, post.Content).Scan(&post_created.Id, &post_created.Title,
		&post_created.Content, &post_created.CreatedAt)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	post_created, errJson := appRep.AddPostCategories(post_created, post.PostCategories)
	if errJson != nil {
		return nil, errJson
	}
	username, errJson := appRep.GetUserNameById(post.UserId)
	if errJson != nil {
		return nil, errJson
	}
	post_created.Username = username
	return post_created, nil
}

// all the posts
// add the offset and the limit after
func (appRep *AppRepository) GetPosts(user_id, offset int) ([]models.Post, *models.ErrorJson) {
	var where string
	if offset == 0 {
		where = ""
	} else {
		where = `WHERE posts.postID < ?`
	}

	var posts []models.Post

	query := fmt.Sprintf(`
	with
    cte_likes as (
        select
            entityID,
            count(*) as total_likes
        from
            reactions
        WHERE
            reactions.entityTypeID = 1
            AND reactions.reaction = 1
        group by
            entityID
    ),
    cte_comments as (
        SELECT
            postID,
            count(*) as total_comments
        from
            comments
        GROUP BY
            postID
    )
	SELECT
		DISTINCT
		users.nickname,
		posts.postID,
		posts.createdAt,
		posts.title,
		posts.content,
		coalesce(cte_likes.total_likes, 0) as total_likes,
		coalesce(cte_comments.total_comments, 0) as total_comments,
		coalesce(reactions.userID,0) as liked
	FROM
		posts
		INNER JOIN users ON posts.userID = users.userID
		LEFT JOIN cte_likes ON posts.postID = cte_likes.entityID
		LEFT JOIN cte_comments ON cte_comments.postID = posts.postID
		LEFT JOIN reactions ON reactions.entityID = posts.postID 
		AND reactions.userID = ? AND reactions.reaction = 1 AND reactions.entityTypeID = 1
	  %v
	ORDER BY
		posts.createdAt DESC
		LIMIT 10;
	`, where)
	rows, err := appRep.db.Query(query, user_id, offset)

	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	if rows.Err() == sql.ErrNoRows {
		return posts, nil
	}

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Username, &post.Id, &post.CreatedAt, &post.Title,
			&post.Content, &post.TotalLikes, &post.TotalComments, &post.Liked); err != nil {
			return posts, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		query_fetch_categories := `
		SELECT  categories.category
		FROM categories INNER JOIN postCategories ON 
		categories.categoryID = postCategories.categoryID
		INNER JOIN posts ON postCategories.postID = posts.postID 
		WHERE posts.postID = ? 
		`
		rows_, errQuery := appRep.db.Query(query_fetch_categories, post.Id)
		if errQuery != nil {
			return posts, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v   5", err)}
		}
		categories := []any{}
		for rows_.Next() {
			var category string
			errScan := rows_.Scan(&category)
			if errScan != nil {
				return posts, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
			}
			categories = append(categories, category)

		}
		post.PostCategories = append(post.PostCategories, categories...)

		posts = append(posts, post)

	}
	defer rows.Close()
	return posts, nil
}

// got everything done here
