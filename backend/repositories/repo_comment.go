package repositories

import (
	"database/sql"
	"fmt"

	"real-time-forum/backend/models"
)

func (appRep *AppRepository) CreateComment(comment *models.Comment) (*models.Comment, *models.ErrorJson) {
	comment_created := &models.Comment{}
	query := `INSERT INTO comments(postID, userID, content)  VALUES(?, ?, ?) 
	RETURNING commentID, content, createdAt;`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	defer stmt.Close()
	if err := stmt.QueryRow(comment.PostId, comment.UserId, comment.Content).Scan(
		&comment_created.Id, &comment_created.Content,
		&comment_created.CreatedAt); err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	username, errJSon := appRep.GetUserNameById(comment.UserId)
	if errJSon != nil {
		return nil, models.NewErrorJson(500, *errJSon)
	}
	comment_created.Username = username
	return comment_created, nil
}

// But hna comments dyal wa7d l post specific
func (appRep *AppRepository) GetComments(user_id, postId, offset int) ([]models.Comment, error) {
	var comments []models.Comment
	query := `
	with
    cte_likes as (
        SELECT
            entityID,
            count(*) as total_likes
        FROM
            reactions
            INNER JOIN types ON reactions.entityTypeID = types.entityTypeID
        WHERE
            types.entityType = "comment"
			AND reactions.reaction = 1
        GROUP BY
            entityID
    )
	SELECT
		users.nickname,
		comments.commentID,
		content,
		comments.createdAt,
		coalesce(cte_likes.total_likes, 0) as total_likes,
        coalesce(reactions.userID,0) as liked
	FROM
		comments
		INNER JOIN users ON comments.userID = users.userID
		LEFT JOIN cte_likes ON cte_likes.entityID = comments.commentID
        LEFT JOIN reactions ON comments.commentID = reactions.entityID 
        AND reactions.userID = ?  AND reactions.reaction =1 AND reactions.entityTypeID = 2
	WHERE
		comments.postID = ?
	ORDER BY
		comments.createdAt DESC
	LIMIT
		10
	OFFSET
		?;
	`
	rows, err := appRep.db.Query(query, user_id, postId, offset)
	if rows.Err() == sql.ErrNoRows {
		return comments, nil
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment models.Comment
		if err = rows.Scan(&comment.Username, &comment.Id, &comment.Content,
			&comment.CreatedAt, &comment.TotalLikes, &comment.Liked); err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	defer rows.Close()
	return comments, nil
}
