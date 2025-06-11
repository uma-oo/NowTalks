package repositories

import (
	"fmt"

	"real-time-forum/backend/models"
)

func (repo *AppRepository) AddMessage(message *models.Message) (*models.Message, *models.ErrorJson) {
	message_created := &models.Message{}
	query := `INSERT INTO messages (senderID,receiverID,message, createdAt) 
	VALUES (?,?,?,?) RETURNING senderID ,receiverID ,message, createdAt;`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	defer stmt.Close()
	if err = stmt.QueryRow(message.SenderID, message.ReceiverID, message.Message, message.CreatedAt).Scan(
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

// the one logged in trying to see the messages will not be got from the query
// sender and receiver and the offset and limit als
func (repo *AppRepository) GetMessages(sender_id, receiver_id, offset int) ([]models.Message, *models.ErrorJson) {
	var messages []models.Message
	query := `
	SELECT
		s.nickname AS sender,
		r.nickname AS receiver,
		messages.message,
		messages.createdAt,
		messages.messageID
	FROM
		messages INNER JOIN users s
		ON messages.senderID = s.userID 
		JOIN users r ON 
		messages.receiverID = r.userID
	WHERE
		senderID IN (?, ?)
		AND receiverID IN (?, ?)
		-- AND messages.messageID < ?
	ORDER BY  messages.createdAt  DESC
	LIMIT
		10
`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	rows, err := stmt.Query(sender_id, receiver_id, sender_id, receiver_id)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.SenderUsername,
			&message.ReceiverUsername, &message.Message,
			&message.CreatedAt, &message.MessageID); err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		messages = append(messages, message)
	}

	return messages, nil

}
