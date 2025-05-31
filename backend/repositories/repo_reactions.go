package repositories

import (
	"database/sql"
	"fmt"

	"real-time-forum/backend/models"
)

//  The post reactions part

func (appRepo *AppRepository) AddReaction(reaction *models.Reaction, type_reaction int) *models.ErrorJson {
	query := `INSERT INTO reactions 
	(entityTypeID, entityID,reaction,userID) VALUES (?,?,?,?)`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	_, err = stmt.Exec(reaction.EntityTypeId, reaction.EntityId, type_reaction, reaction.UserId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

func (appRepo *AppRepository) UpdateReactionLike(reaction *models.Reaction) *models.ErrorJson {
	query := `UPDATE reactions SET reaction = CASE reaction
              WHEN 0 THEN 1
			  WHEN -1 THEN 1
              ELSE 0
              END
	          WHERE reactionID = ? ;`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v h", err)}
	}
	defer stmt.Close()
	_, err = stmt.Exec(reaction.Id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v hh", err)}
	}

	return nil
}

func (appRepo *AppRepository) UpdateReactionDislike(reaction *models.Reaction) *models.ErrorJson {
	query := `UPDATE reactions SET reaction = CASE reaction
              WHEN 0 THEN -1
			  WHEN 1 THEN -1
              ELSE 0
              END
	          WHERE reactionID = ? ;`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v h", err)}
	}
	defer stmt.Close()
	_, err = stmt.Exec(reaction.Id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v hh", err)}
	}

	return nil
}





func (appRepo *AppRepository) HanldeReaction(reaction *models.Reaction) (*models.Reaction, *models.ErrorJson) {
	reaction_existed := &models.Reaction{}
	query := `SELECT * FROM reactions WHERE userID = ? AND entityTypeID = ? AND entityID = ?`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v jj", err)}
	}
	if err := stmt.QueryRow(reaction.UserId, reaction.EntityTypeId, reaction.EntityId).Scan(
		&reaction_existed.Id, &reaction_existed.EntityTypeId,
		&reaction_existed.EntityId, &reaction_existed.Reaction, &reaction_existed.UserId); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v jjj", err)}
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

func (appRepo *AppRepository) CheckEntityID(reaction *models.Reaction, type_entity string) *models.ErrorJson {
	var query string
	var entity int
	switch type_entity {
	case "comment":
		query = `SELECT exists(SELECT 1 FROM comments WHERE commentID = ? );`
		if err := appRepo.db.QueryRow(query, reaction.EntityId).Scan(&entity); err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v lhiih", err)}
		}

	case "post":
		query = `SELECT exists(SELECT 1 FROM posts WHERE postID = ? );`
		if err := appRepo.db.QueryRow(query, reaction.EntityId).Scan(&entity); err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}

	}
     // exists will return 0 if there is no row thar matches dakshi li 3ndna 
	if entity == 0 {
		return &models.ErrorJson{Status: 400, Message: &models.ReactionErr{
			EntityId: "ERROR!! Wrong EntityID field!",
		}}
	}
	return nil
}
