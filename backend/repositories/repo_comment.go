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
