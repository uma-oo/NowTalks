package repositories

import (
	"database/sql"
	"fmt"

	"real-time-forum/backend/models"
)

//  The post reactions part

func (appRepo *AppRepository) AddReactionLike(reaction *models.Reaction) *models.ErrorJson {
	query := `INSERT INTO reactions 
	(entityTypeID, entityID,reaction,userID) VALUES (?,?,?,?)`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	_, err = stmt.Exec(reaction.EntityTypeId, reaction.EntityId, 1, reaction.UserId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

func (appRepo *AppRepository) UpdateReactionLike(reaction *models.Reaction) *models.ErrorJson {
	fmt.Println("hnaaaaaaaaaaaaaaaaaa")
	query := `UPDATE reactions SET reaction = CASE reaction
              WHEN 0 THEN 1
              ELSE 0
              END
	          WHERE reactionID = ? ;`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		fmt.Println("err hhh", err)
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	fmt.Println("err  ddd",reaction.Id )
	_, err = stmt.Exec(reaction.Id)
	if err != nil {
		fmt.Println("err  ddd", err)
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return nil
}

func (appRepo *AppRepository) HanldeReaction(reaction *models.Reaction) (*models.Reaction, *models.ErrorJson) {
	reaction_existed := &models.Reaction{}
	query := `SELECT * FROM reactions WHERE userID = ? AND entityTypeID = ? AND entityID = ?`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	if err := stmt.QueryRow(reaction.UserId, reaction.EntityTypeId, reaction.EntityId).Scan(
		&reaction_existed.Id, &reaction_existed.EntityTypeId,
		&reaction_existed.EntityId, &reaction_existed.Reaction, &reaction_existed.UserId); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return reaction_existed, nil
}

func (appRepo *AppRepository) GetTypeIdByName(type_entity string) int {
	var id int
	query := `SELECT entityTypeID FROM types WHERE entityType = ?`
	if err := appRepo.db.QueryRow(query, type_entity).Scan(&id); err != nil {
		return 0
	}
	return id
}

func (appRepo *AppRepository) CheckEntityID(type_entity string, reaction *models.Reaction) *models.ErrorJson {
	var entity any
	var query string
	switch type_entity {
	case "comment":
		query = `SELECT * FROM comments WHERE commentID = ? ;`
		v, ok := entity.(models.Comment)
		if ok {
			if err := appRepo.db.QueryRow(query, reaction.EntityId).Scan(&v); err != nil {
				return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v lhiih", err)}
			}
		}
	case "post":
		query = `SELECT * FROM posts WHERE postID = ? ;`
		v, ok := entity.(models.Post)
		if ok {
			if err := appRepo.db.QueryRow(query, reaction.EntityId).Scan(&v); err != nil {
				return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v hnaaa", err)}
			}
		}

	}

	return nil
}
