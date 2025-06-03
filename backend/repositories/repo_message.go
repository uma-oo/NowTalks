package repositories

import (
	"fmt"
	"real-time-forum/backend/models"
)



func (repo *AppRepository) AddMessage(message *models.Message) (*models.Message, *models.ErrorJson) {
	
	message_created := &models.Message{}
	query := `INSERT INTO messages (senderID,receiverID,message) 
	VALUES (?,?,?) RETURNING senderID ,receiverID ,message, createdAt;`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	defer stmt.Close()
	if err = stmt.QueryRow( message.SenderID, message.ReceiverID, message.Message).Scan(
		&message_created.SenderID, &message_created.ReceiverID,
		&message_created.Message, &message_created.CreatedAt); err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v 1", err))
	}
	receiver, errRec := repo.GetUserNameById(message.ReceiverID)
	sender, errSen := repo.GetUserNameById(message.SenderID)
	if errRec != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v 2", errRec))
	}
	if errSen != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v 3", errSen))
	}
	message_created.ReceiverUsername = receiver
	message_created.SenderUsername = sender
	return message_created, nil
}
