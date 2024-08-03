package dto

type MessageDto struct {
	Id      int    `json:"id"`
	ChatId  int    `json:"chat_id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

