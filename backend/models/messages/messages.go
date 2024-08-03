package model

type Message struct {
	Id      int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	UserId  int    `gorm:"NOT NULL"`
	ChatId  int    `gorm:"NOT NULL"`
	Content string `gorm:"type:varchar(500);NOT NULL"`
}

type Messages []Message
