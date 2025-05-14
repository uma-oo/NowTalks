package repositories

import (
	"real-time-forum/backend/models"
)

func (appRep *AppRepository) CreateComment(comment *models.Comment) error {
	query := `INSERT INTO comments(postID, userID, content)  VALUES(?, ?, ?)`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
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

// To write the data back f response we need to have (userId , postId) Very known
// * userId to recuperate from the session ??? (user is authenticated)
// * postId tp recuperate from the query

// BUT THE COMBINAISON OF postID+userID is not unique !!
// the latest one
// 204 no content 
func (appRepo *AppRepository) GetWrittenComment(userId int, postId int) (*models.Comment, *models.ErrorJson) {
	comment := models.NewComment()
	query := `SELECT commentID, userID, postID, createdAt , content FROM comments WHERE userID = ? AND postID = ?`
	if err := appRepo.db.QueryRow(query, userId, postId).Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.CreatedAt, &comment.Content); err != nil {
		return nil, models.NewErrorJson(204, models.Comment{})
	}

	return comment, nil
}


