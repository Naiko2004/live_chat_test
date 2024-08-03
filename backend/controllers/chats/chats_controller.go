package chats

import (
	"backend_chat/dto"
	e "backend_chat/errors"
	chatsService "backend_chat/services/chats"
	usersService "backend_chat/services/users"
	"strconv"

	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateChat(c *gin.Context) {
	var chat dto.ChatCreateDto
	if err := c.ShouldBindJSON(&chat); err != nil {
		log.Printf("Error binding chat: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	e := chatsService.CreateChat(chat)
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func GetChatsByUserId(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	userId, err := usersService.ValidateToken(token)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	chats, e := chatsService.GetChatsByUserId(userId)

	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	c.JSON(http.StatusOK, chats)
}

func JoinChat(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	userId, e := usersService.ValidateToken(token)
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	chatId, err := strconv.Atoi(c.Param("chatId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat id"})
		return
	}

	e = chatsService.JoinChat(userId, chatId)

	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "joined"})
}

func LeaveChat(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	userId, e := usersService.ValidateToken(token)
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	chatId, err := strconv.Atoi(c.Param("chatId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat id"})
		return
	}

	e = chatsService.LeaveChat(userId, chatId)

	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "left"})
}