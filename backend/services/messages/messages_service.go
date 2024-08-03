package messages

import (
	chatsClient "backend_chat/clients/chats"
	messagesClient "backend_chat/clients/messages"
	usersClient "backend_chat/clients/users"
	"backend_chat/dto"
	messagesModel "backend_chat/models/messages"
	usersService "backend_chat/services/users"

	e "backend_chat/errors"
)

func CreateMessage(messageDto dto.MessageCreateDto) e.ApiError {
	message := messagesModel.Message{
		ChatId:  messageDto.ChatId,
		UserId:  messageDto.UserId,
		Content: messageDto.Text,
	}
	return messagesClient.CreateMessage(message)
}

func CreateMessageWS(messageDto dto.MessageCreateWSDto) e.ApiError {

	userId, err := usersService.ValidateToken(messageDto.Token)

	if err != nil {
		return err
	}

	message := messagesModel.Message{
		ChatId:  messageDto.ChatId,
		UserId:  userId,
		Content: messageDto.Text,
	}
	return messagesClient.CreateMessage(message)
}

func GetMessagesByChatId(chatId int) ([]dto.MessageDto, e.ApiError) {
	messagesDb, err := messagesClient.GetMessagesByChatId(chatId)
	if err != nil {
		return []dto.MessageDto{}, err
	}

	var messagesDto []dto.MessageDto
	for _, message := range messagesDb {
		messageDto := dto.MessageDto{
			Id:      message.Id,
			ChatId:  message.ChatId,
			UserId:  message.UserId,
			Content: message.Content,
		}
		messagesDto = append(messagesDto, messageDto)
	}

	return messagesDto, nil
}

func GetAllMessageById(messageId int) (dto.MessageAllDto, e.ApiError) {
	messageDb, err := messagesClient.GetMessageById(messageId)
	if err != nil {
		return dto.MessageAllDto{}, err
	}

	chatDb, err := chatsClient.GetChatById(messageDb.ChatId)
	if err != nil {
		return dto.MessageAllDto{}, err
	}

	userDb, err := usersClient.GetUserById(messageDb.UserId)
	if err != nil {
		return dto.MessageAllDto{}, err
	}

	messageDto := dto.MessageAllDto{
		Id:       messageDb.Id,
		ChatId:   messageDb.ChatId,
		ChatName: chatDb.Name,
		UserId:   messageDb.UserId,
		Username: userDb.Username,
		Content:  messageDb.Content,
	}

	return messageDto, nil
}
