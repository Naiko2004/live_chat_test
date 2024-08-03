package messages

import (
	e "backend_chat/errors"
	messagesModel "backend_chat/models/messages"
	"log"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func CreateMessage(message messagesModel.Message) e.ApiError {
	err := Db.Create(&message).Error

	if err != nil {
		log.Printf("Error creating message: %v", err)
		return e.NewApiError("Error creating message", "internal_server_error", 500, e.CauseList{})
	}

	return nil
}

func GetMessagesByChatId(chatId int) ([]messagesModel.Message, e.ApiError) {
	var messages []messagesModel.Message
	err := Db.Where("chat_id = ?", chatId).Find(&messages).Error

	if err != nil {
		log.Printf("Error getting messages by chat id: %v", err)
		return []messagesModel.Message{}, e.NewNotFoundApiError("Messages not found")
	}

	return messages, nil
}

func GetMessageById(messageId int) (messagesModel.Message, e.ApiError) {
	var message messagesModel.Message
	err := Db.Where("id = ?", messageId).First(&message).Error

	if err != nil {
		log.Printf("Error getting message by id: %v", err)
		return messagesModel.Message{}, e.NewNotFoundApiError("Message not found")
	}

	return message, nil
}
