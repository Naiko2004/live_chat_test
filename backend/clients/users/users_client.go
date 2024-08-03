package users

import (
	e "backend_chat/errors"
	userModel "backend_chat/models/users"
	"log"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetUserById(id int) (userModel.User, e.ApiError) {
	var user userModel.User
	err := Db.First(&user, id).Error

	if err != nil {
		log.Printf("Error getting user by id: %v", err)
		return userModel.User{}, e.NewNotFoundApiError("User not found")
	}

	return user, nil
}

func GetUserByUsername(username string) (userModel.User, e.ApiError) {
	var user userModel.User
	err := Db.Where("username = ?", username).First(&user).Error

	if err != nil {
		log.Printf("Error getting user by username: %v", err)
		return userModel.User{}, e.NewNotFoundApiError("User not found")
	}

	return user, nil
}

func CreateUser(user userModel.User) e.ApiError {
	err := Db.Create(&user).Error

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return e.NewApiError("Error creating user", "internal_server_error", 500, e.CauseList{})
	}

	return nil
}

func CheckUsername(username string) (bool, e.ApiError) {
	var user userModel.User
	err := Db.Where("username = ?", username).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		log.Printf("Error checking username: %v", err)
		return false, e.NewApiError("Error checking username", "internal_server_error", 500, e.CauseList{})
	}

	return true, nil
}
