package dto

type MessageCreateWSDto struct {
	ChatId int    `json:"chat_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
	Text   string `json:"text" binding:"required"`
}
