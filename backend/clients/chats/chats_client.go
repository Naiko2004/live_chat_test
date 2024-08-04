package chats

import (
	e "backend_chat/errors"
	chatModel "backend_chat/models/chats"
	"log"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetChatById(chatId int) (chatModel.Chat, e.ApiError) {
	var chat chatModel.Chat
	err := Db.Where("id = ?", chatId).First(&chat).Error
	if err != nil {
		log.Printf("Error getting chat by id: %v", err)
		return chatModel.Chat{}, e.NewNotFoundApiError("Chat not found")
	}
	return chat, nil
}

func CreateChat(chat *chatModel.Chat) e.ApiError {
	err := Db.Create(&chat).Error

	if err != nil {
		log.Printf("Error creating chat: %v", err)
		return e.NewApiError("Error creating chat", "internal_server_error", 500, e.CauseList{})
	}

	return nil
}

func CreateChatParticipant(chatUser chatModel.ChatParticipant) e.ApiError {
	err := Db.Create(&chatUser).Error
	if err != nil {
		log.Printf("Error creating chat user: %v", err)
		return e.NewApiError("Error creating chat user", "internal_server_error", 500, e.CauseList{})
	}
	return nil
}

func DeleteChatParticipant(chatUser chatModel.ChatParticipant) e.ApiError {
	err := Db.Where("chat_id = ? AND user_id = ?", chatUser.ChatId, chatUser.UserId).Delete(&chatUser).Error
	if err != nil {
		log.Printf("Error deleting chat user: %v", err)
		return e.NewApiError("Error deleting chat user", "internal_server_error", 500, e.CauseList{})
	}
	return nil
}

func GetChatsByUserId(userId int) ([]chatModel.Chat, e.ApiError) {
	var chats []chatModel.Chat
	err := Db.Table("chats").Joins("JOIN chat_participants ON chats.id = chat_participants.chat_id").Where("chat_participants.user_id = ?", userId).Find(&chats).Error
	if err != nil {
		log.Printf("Error getting chats by user id: %v", err)
		return nil, e.NewApiError("Error getting chats by user id", "internal_server_error", 500, e.CauseList{})
	}
	return chats, nil
}

func GetChats() ([]chatModel.Chat, e.ApiError) {
	var chats []chatModel.Chat
	err := Db.Find(&chats).Error
	if err != nil {
		log.Printf("Error getting chats: %v", err)
		return nil, e.NewApiError("Error getting chats", "internal_server_error", 500, e.CauseList{})
	}
	return chats, nil
}
