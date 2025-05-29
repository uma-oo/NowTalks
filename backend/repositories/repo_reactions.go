package repositories

import (
	"fmt"

	"real-time-forum/backend/models"
)

//  The post reactions part

func (appRepo *AppRepository) AddReaction(reaction *models.Reaction) *models.ErrorJson {
	query := `INSERT INTO reactions 
	(entityTypeID, entityID,reaction,userID) VALUES (?,?,?,?)`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	_, err = stmt.Exec(reaction.EntityTypeId, reaction.EntityId, reaction.Reaction, reaction.UserId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

func (appRepo *AppRepository) UpdateReaction(reaction *models.Reaction, new_reaction int) *models.ErrorJson {
	query := `UPDATE reactions SET reaction = ? WHERE reactionID = ?`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	_, err = stmt.Exec(new_reaction, reaction.Id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return nil
}
