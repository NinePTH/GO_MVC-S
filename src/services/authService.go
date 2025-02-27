package services

import (
	"errors"

	"github.com/NinePTH/GO_MVC-S/src/middlewares"
	"github.com/NinePTH/GO_MVC-S/src/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username string, password string) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	data := map[string]interface{}{
		"username": username,
		"password": string(hashedPassword),
	}

	rowsAffected, err := InsertData("users", data)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func AuthenticateUser(username string, password string) (*models.Token, error) {
	fields := []string{"id", "username", "password"}
	whereCondition := "username =$1"
	whereArgs := []interface{}{username}

	result, err := SelectData("users", fields, true, whereCondition, whereArgs,"id")
	if err != nil {
		return nil, err
	}

	if len(result) == 0{
		return nil, errors.New("user not found")
	}

	user := result[0]
	storedPassword := user["password"].(string)

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := middlewares.GenerateJWT(username)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &models.Token{
		Token: token,
	}, nil
}