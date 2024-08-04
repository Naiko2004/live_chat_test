package users

import (
	"backend_chat/dto"
	usersService "backend_chat/services/users"

	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user dto.UserMinDto
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	e := usersService.CreateUser(user)
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	user, e := usersService.GetUserById(id)
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var user dto.UserMinDto
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	token, e := usersService.Login(user)
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetMe(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token := strings.Split(authHeader, " ")[1]

	userId, e := usersService.ValidateToken(token)
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	user, e := usersService.GetUserById(userId)
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}
	c.JSON(http.StatusOK, user)
}
