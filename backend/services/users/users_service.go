package users

import (
	usersClient "backend_chat/clients/users"
	"backend_chat/dto"
	usersMoodel "backend_chat/models/users"

	e "backend_chat/errors"

	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, e.ApiError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", e.NewApiError("Error hashing password", "internal_server_error", 500, e.CauseList{})
	}

	return string(hashedPassword), nil
}

func CreateUser(user dto.UserMinDto) e.ApiError {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	usernameExists, err := usersClient.CheckUsername(user.Username)

	if err != nil {
		return err
	}

	if usernameExists {
		return e.NewApiError("Username already exists", "bad_request", 400, e.CauseList{})
	}

	userModel := usersMoodel.User{
		Username: user.Username,
		Password: string(hashedPassword),
	}

	err = usersClient.CreateUser(userModel)
	if err != nil {
		return err
	}

	return nil
}

func GetUserById(id int) (dto.UserMinDto, e.ApiError) {
	userDB, err := usersClient.GetUserById(id)
	if err != nil {
		return dto.UserMinDto{}, err
	}

	user := dto.UserMinDto{
		Username: userDB.Username,
	}

	return user, nil
}

func createToken(id int) (string, e.ApiError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte("Meltryllis"))
	if err != nil {
		return "", e.NewApiError("Error creating token", "internal_server_error", 500, e.CauseList{})
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (int, e.ApiError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("Meltryllis"), nil
	})

	if err != nil {
		return 0, e.NewUnauthorizedApiError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, e.NewUnauthorizedApiError("Invalid token")
	}

	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		return 0, e.NewUnauthorizedApiError("Token has expired")
	}

	id := int(claims["id"].(float64))

	return id, nil
}
func Login(user dto.UserMinDto) (string, e.ApiError) {
	userDB, err := usersClient.GetUserByUsername(user.Username)
	if err != nil {
		return "", err
	}

	err1 := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err1 != nil {
		return "", e.NewApiError("Invalid password", "bad_request", 400, e.CauseList{})
	}

	token, err := createToken(userDB.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}
