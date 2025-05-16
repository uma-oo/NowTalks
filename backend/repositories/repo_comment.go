package repositories

import (
	"fmt"

	"real-time-forum/backend/models"
)

func (appRep *AppRepository) CreateComment(comment *models.Comment) (*models.Comment, *models.ErrorJson) {
	comment_created := &models.Comment{}
	query := `INSERT INTO comments(postID, userID, content)  VALUES(?, ?, ?) RETURNING *`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	defer stmt.Close()
	if err := stmt.QueryRow(comment.PostId, comment.UserId, comment.Content).Scan(
		&comment_created.Id, &comment_created.PostId,
		&comment_created.UserId, &comment_created.CreatedAt,
		&comment_created.Content); err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	username, errJSon := appRep.getUserNameById(comment_created.UserId)
	if errJSon != nil {
		return nil, models.NewErrorJson(500, *errJSon)
	}
	comment_created.Username = username
	return comment_created, nil
}

// But hna comments dyal wa7d l post specific
func (appRep *AppRepository) GetComments(postId int, limit int, offset int) ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT commentID, userID, postID, createdAt , content 
	FROM comments 
	WHERE postID = ?
	ORDER BY createdAt
	LIMIT ? OFFSET ? ;`
	rows, err := appRep.db.Query(query, postId, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var comment models.Comment
		if err = rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.CreatedAt, &comment.Content); err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	defer rows.Close()
	return comments, nil
}





