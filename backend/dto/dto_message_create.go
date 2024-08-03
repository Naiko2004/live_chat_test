package dto

type MessageCreateDto struct {
	ChatId int    `json:"chat_id" binding:"required"`
	UserId int    `json:"user_id" binding:"required"`
	Text   string `json:"text" binding:"required"`
}
