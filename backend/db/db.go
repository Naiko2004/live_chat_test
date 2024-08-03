package db

import (
	// Clients
	chatsClient "backend_chat/clients/chats"
	messagesClient "backend_chat/clients/messages"
	userClient "backend_chat/clients/users"

	// Models
	chatModel "backend_chat/models/chats"
	messageModel "backend_chat/models/messages"
	userModel "backend_chat/models/users"

	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB Connections Parameters
	DBName := os.Getenv("CHAT_DB_NAME")
	DBUser := os.Getenv("CHAT_DB_USER")
	DBPass := os.Getenv("CHAT_DB_PASS")
	DBHost := os.Getenv("CHAT_DB_HOST")
	DBPort := os.Getenv("CHAT_DB_PORT")

	if DBName == "" || DBUser == "" || DBHost == "" || DBPort == "" {
		log.Fatal("Database connection parameters missing")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", DBUser, DBPass, DBHost, DBPort, DBName)
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		log.Fatal(err)
	} else {
		log.Println("Database connection successful")
		log.Println("Database connection parameters: ", dsn)
	}

	// We add the clients
	userClient.Db = db
	messagesClient.Db = db
	chatsClient.Db = db
}

func StartDbEngine() {
	// Migrate the schema
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&messageModel.Message{})
	db.AutoMigrate(&chatModel.Chat{})
	db.AutoMigrate(&chatModel.ChatParticipant{})

	log.Println("Finishing Migration Database Tables")
}
