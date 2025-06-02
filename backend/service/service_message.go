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
	if trimmedMsg == "" {
		errMessage.Message = "ERROR!! Empty Message Body"
	}
	if len(trimmedMsg) > 1000 {
		errMessage.Message = "ERROR!! Message Body Too Large!"
	}
	if username, _ := service.repo.GetUserNameById(message.ReceiverID); username == "" {
		errMessage.ReceiverID = "ERROR!! The Receiver Specified Does Not Exist!!"
	}
	if errMessage.Message != "" || errMessage.ReceiverID != "" {
		return nil, &models.ErrorJson{Status: 400, Message: errMessage}
	}
	// We can go on and insert the message in the database
	message_created, err := service.repo.AddMessage(message)
	if err != nil {
		return nil, err
	}
	return message_created, nil
}
