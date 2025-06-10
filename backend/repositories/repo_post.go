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
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v hhh", err)}
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.UserId, post.Title, post.Content).Scan(&post_created.Id, &post_created.Title,
		&post_created.Content, &post_created.CreatedAt)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v dddd  ", err)}
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
func (appRep *AppRepository) GetPosts(offset int) ([]models.Post, *models.ErrorJson) {
	var posts []models.Post
	query := `
	with
    cte_likes as (
        select
            entityID,
            count(*) as total_likes
        from
            reactions
            INNER JOIN types ON reactions.entityTypeID = types.entityTypeID
        WHERE
            types.entityType = "post"
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
    users.nickname,
    posts.postID,
    posts.createdAt,
    posts.title,
    posts.content,
    coalesce(cte_likes.total_likes, 0) as total_likes,
    coalesce(cte_comments.total_comments, 0) as total_comments
FROM
    posts
    INNER JOIN users ON posts.userID = users.userID
    LEFT JOIN cte_likes ON posts.postID = cte_likes.entityID
    LEFT JOIN cte_comments ON cte_comments.postID = posts.postID
ORDER BY
    posts.createdAt DESC
	LIMIT 10 offset ?;
	`
	rows, err := appRep.db.Query(query, offset)
	if rows.Err() == sql.ErrNoRows {
		return posts, nil
	}
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Username, &post.Id, &post.CreatedAt, &post.Title, &post.Content, &post.TotalLikes, &post.TotalComments); err != nil {
			return posts, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}

		posts = append(posts, post)

	}
	defer rows.Close()
	return posts, nil
}

// got everything done here
