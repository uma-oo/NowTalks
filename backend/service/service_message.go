package service

import (
	"strings"

	"real-time-forum/backend/models"
)

// let's check wash message huwa hadak
// no need to check the sender Id kayn middleware
func (service *AppService) ValidateMessage(message *models.Message) (*models.Message, *models.ErrorJson) {
	errMessage := models.NewMessageErr()
	trimmedMsg := strings.TrimSpace(message.Message)
	type_message := strings.ToLower(strings.TrimSpace(message.Type))

	if type_message != "message" && type_message != "status" && type_message != "typing" {
		errMessage.Type = "ERROR!! Wrong type of message"
	}

	if trimmedMsg == "" {
		errMessage.Message = "ERROR!! Empty Message Body"
	}
	if len(trimmedMsg) > 1000 {
		errMessage.Message = "ERROR!! Message Body Too Large!"
	}
	if username, _ := service.repo.GetUserNameById(message.ReceiverID); username == "" {
		errMessage.ReceiverID = "ERROR!! The Receiver Specified Does Not Exist!!"
	}

	if message.CreatedAt.IsZero() {
		errMessage.CreatedAt = "ERROR!! The date is not set up!"
	}

	if errMessage.Message != "" || errMessage.ReceiverID != "" || errMessage.Type != "" || errMessage.CreatedAt != "" {
		return nil, &models.ErrorJson{Status: 400, Message: errMessage}
	}

	// We can go on and insert the message in the database
	switch strings.ToLower(message.Type) {
	case "message":
		message_created, err := service.repo.AddMessage(message)
		if err != nil {
			return nil, err
		}
		message_created.Type = type_message
		return message_created, nil
	case "status":

	case "typing":

	}

	// so in this case we only need to update the database (the message exists already)
	// if message.Type == "read" {
	// }
	return nil, nil
}

// from the unread to the read status
func (service *AppService) EditReadStatus(sender_id, receiver_id int) *models.ErrorJson {
	if err := service.repo.EditReadStatus(sender_id, receiver_id); err != nil {
		return err
	}
	return nil
}

func (service *AppService) GetMessages(sender_id, receiver_id, offset int, type_ string) ([]models.Message, *models.ErrorJson) {
	messages, errJson := service.repo.GetMessages(sender_id, receiver_id, offset, type_)
	if errJson != nil {
		return nil, errJson
	}
	return messages, nil
}
