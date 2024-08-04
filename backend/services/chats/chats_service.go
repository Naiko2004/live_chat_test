package chats

import (
	chatsClient "backend_chat/clients/chats"
	usersClient "backend_chat/clients/users"
	"backend_chat/dto"
	chatsModel "backend_chat/models/chats"
	"log"

	e "backend_chat/errors"
)

func CreateChat(chatDto dto.ChatCreateDto) e.ApiError {
	chat := chatsModel.Chat{
		Name: chatDto.Name,
	}
	err := chatsClient.CreateChat(&chat)
	if err != nil {
		return err
	}

	log.Printf("Chat created, name: %s and id: %v ", chat.Name, chat.Id)

	for _, username := range chatDto.Usernames {
		user, err := usersClient.GetUserByUsername(username)

		if err != nil {
			return err
		}
		userId := user.Id
		err = chatsClient.CreateChatParticipant(chatsModel.ChatParticipant{
			ChatId: chat.Id,
			UserId: userId,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func GetChatsByUserId(userId int) ([]dto.ChatMinDtop, e.ApiError) {
	chats, err := chatsClient.GetChatsByUserId(userId)
	if err != nil {
		return nil, err
	}

	var chatsDto []dto.ChatMinDtop
	for _, chat := range chats {
		chatsDto = append(chatsDto, dto.ChatMinDtop{
			Id:   chat.Id,
			Name: chat.Name,
		})
	}

	return chatsDto, nil
}

func JoinChat(userId int, chatId int) e.ApiError {
	err := chatsClient.CreateChatParticipant(chatsModel.ChatParticipant{
		ChatId: chatId,
		UserId: userId,
	})
	if err != nil {
		return err
	}

	return nil
}

func LeaveChat(userId int, chatId int) e.ApiError {
	log.Printf("User %v leaving chat %v", userId, chatId)
	err := chatsClient.DeleteChatParticipant(chatsModel.ChatParticipant{
		ChatId: chatId,
		UserId: userId,
	})
	if err != nil {
		return err
	}

	return nil
}

func GetChats() ([]dto.ChatMinDtop, e.ApiError) {
	chats, err := chatsClient.GetChats()
	if err != nil {
		return nil, err
	}

	var chatsDto []dto.ChatMinDtop
	for _, chat := range chats {
		chatsDto = append(chatsDto, dto.ChatMinDtop{
			Id:   chat.Id,
			Name: chat.Name,
		})
	}

	return chatsDto, nil
}