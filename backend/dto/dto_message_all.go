package dto

type MessageAllDto struct {
	Id      int    `json:"id"`
	ChatId  int    `json:"chat_id"`
	ChatName string `json:"chat_name"`
	UserId  int    `json:"user_id"`
	Username string `json:"user_name"`
	Content string `json:"content"`
}

type MessagesAllDto []MessageAllDto