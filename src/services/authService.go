package services

import "golang.org/x/crypto/bcrypt"

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