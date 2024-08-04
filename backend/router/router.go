package router

import (
	"backend_chat/controllers/messages"
	"backend_chat/controllers/users"
	"backend_chat/controllers/chats"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	// Websocket
	// Messages
	engine.GET("/ws/chat/:chatId", messages.HandleConnections)

	// Public
	// Users
	engine.POST("/users", users.CreateUser)
	engine.GET("/users/:id", users.GetUserById)
	engine.PUT("/login", users.Login)

	// Chats
	engine.GET("/chats/all", chats.GetChats)

	// Private (requires token)
	// Messages
	engine.POST("/messages/:chatId", messages.CreateMessage)
	engine.GET("/messages/:chatId", messages.GetMessagesByChatId)
	engine.GET("message/:id/all", messages.GetAllMessageById)

	// Chats
	engine.POST("/chats", chats.CreateChat)
	engine.GET("user/chats", chats.GetChatsByUserId)
	engine.POST("/chats/:chatId", chats.JoinChat)
	engine.DELETE("/chats/:chatId", chats.LeaveChat) // Hard delete

	// Users
	engine.GET("/me", users.GetMe)
}
