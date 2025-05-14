package repositories

import (
	"real-time-forum/backend/models"
)

func (appRep *AppRepository) CreateComment(comment *models.Comment) error {
	query := `INSERT INTO comments(postID, userID, content)  VALUES(?, ?, ?)`
	stmt, err := appRep.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(comment.PostId, comment.UserId, comment.Content); err != nil {
		return err
	}
	return nil
}

// But hna comments dyal wa7d l post specific
func (appRep *AppRepository) GetComments(postId int) ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT commentID, userID, postID, createdAt , content FROM comments WHERE postID = ?`
	rows, err := appRep.db.Query(query, postId)
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
