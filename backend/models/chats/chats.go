package model

type Chat struct {
	Id   int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(100);NOT NULL"`
}

type Chats []Chat

type ChatParticipant struct {
	ChatId int `gorm:"type:INT;NOT NULL"`
	UserId int `gorm:"type:INT;NOT NULL"`
}

type ChatParticipants []ChatParticipant
