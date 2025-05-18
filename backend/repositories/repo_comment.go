package repositories

import (
	"database/sql"
	"fmt"

	"real-time-forum/backend/models"
)

func (appRep *AppRepository) CreateComment(comment *models.Comment) (*models.Comment, *models.ErrorJson) {
	comment_created := &models.Comment{}
	query := `INSERT INTO comments(postID, userID, content)  VALUES(?, ?, ?) RETURNING commentID, content, createdAt `
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
	username, errJSon := appRep.getUserNameById(comment.UserId)
	if errJSon != nil {
		return nil, models.NewErrorJson(500, *errJSon)
	}
	comment_created.Username = username
	return comment_created, nil
}

// But hna comments dyal wa7d l post specific
func (appRep *AppRepository) GetComments(postId int, limit int, offset int) ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT users.nickname,commentID, postID, createdAt , content 
	FROM comments INNER JOIN on users.userID = comments.userID
	WHERE postID = ?
	ORDER BY createdAt DESC
	LIMIT ? OFFSET ? ;`
	rows, err := appRep.db.Query(query, postId, limit, offset)
	if rows.Err() == sql.ErrNoRows {
		return comments, nil
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment models.Comment
		if err = rows.Scan(&comment.Username, &comment.Id, &comment.PostId, &comment.CreatedAt, &comment.Content); err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	defer rows.Close()
	return comments, nil
}
