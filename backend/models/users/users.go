package model

type User struct {
	Id       int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Username string `gorm:"type:varchar(100);NOT NULL"`
	Password string `gorm:"type:varchar(100);NOT NULL"`
}

type Users []User
