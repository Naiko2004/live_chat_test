package messages

import (
	"log"
	"net/http"

	"encoding/json"
	"strconv"
	"strings"

	"backend_chat/dto"
	e "backend_chat/errors"

	messageService "backend_chat/services/messages"
	usersService "backend_chat/services/users"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(c *gin.Context) {
	w := c.Writer
	r := c.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		_, msgJson, readErr := conn.ReadMessage()
		if readErr != nil {
			log.Println("Error reading message: ", readErr)
			return
		}
		log.Println("Message received: ", string(msgJson))
		// Process the Json into a DTO
		var messageDto dto.MessageCreateWSDto
		unmarshalErr := json.Unmarshal(msgJson, &messageDto)
		if unmarshalErr != nil {
			log.Println("Error unmarshalling: ", unmarshalErr)
			return
		}
		// Create msg in db
		createErr := messageService.CreateMessageWS(messageDto)
		if createErr != nil {
			log.Println("Error creating message: ", createErr)
			return
		}
		// Send the message to all clients
		writeErr := conn.WriteJSON(messageDto)
		if writeErr != nil {
			log.Println("Error sending the msg: ", writeErr)
			return
		}
		log.Println("Message received and sent: ", messageDto)
	}
}

func CreateMessage(c *gin.Context) {
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

	chatId, err1 := strconv.Atoi(c.Param("chatId"))
	if err1 != nil {
		log.Printf("Error converting chat id to int: %v", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat id"})
		return
	}

	var messageDto dto.MessageCreateDto
	messageDto.ChatId = chatId
	messageDto.UserId = userId
	if err := c.ShouldBindJSON(&messageDto); err != nil {
		log.Printf("Error binding message: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = messageService.CreateMessage(messageDto)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Message created"})
}

func GetMessagesByChatId(c *gin.Context) {
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

	_, err := usersService.ValidateToken(token)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	chatId, err1 := strconv.Atoi(c.Param("chatId"))
	if err1 != nil {
		log.Printf("Error converting chat id to int: %v", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat id"})
		return
	}

	messages, err := messageService.GetMessagesByChatId(chatId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func GetAllMessageById(c *gin.Context) {
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

	_, err := usersService.ValidateToken(token)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	messageId, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		log.Printf("Error converting message id to int: %v", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message id"})
		return
	}

	message, err := messageService.GetAllMessageById(messageId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
