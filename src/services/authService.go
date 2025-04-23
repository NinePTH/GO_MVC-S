package services

import (
	"errors"
	"fmt"

	"github.com/NinePTH/GO_MVC-S/src/middlewares"
	"github.com/NinePTH/GO_MVC-S/src/models/auth"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username string, password string, role string, id string) (int64, error) {

	table := ""
	fields := []string{"*"}
	whereCondition := ""

	if role == "patient" {
		table = "Patient"
		whereCondition = "patient_id = $1 AND user_id IS NULL"
	} else if role == "HR" || role == "medical_personnel" {
		table = "Employee"
		whereCondition = "employee_id = $1 AND user_id IS NULL"
	} else {
		return 0, errors.New("Invalid role")
	}

	result, err := SelectData(table, fields, true, whereCondition, []interface{}{id}, false, "", "")

	if err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, fmt.Errorf("There is no patient with id %s or the patient has already been registered", id)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	data := map[string]interface{}{
		"username": username,
		"password": string(hashedPassword),
		"role":     role,
	}

	insertResult, err := InsertData("users", data)
	if err != nil {
		return 0, err
	}

	if insertResult == 0 {
		return 0, errors.New("Failed to insert user")
	}

	// ทำให้มัน อัพเดต user_id ใน patient table

	fields = []string{"user_id", "username"}
	whereCondition = "username =$1"
	whereArgs := []interface{}{username}

	result, err = SelectData("users", fields, true, whereCondition, whereArgs, false, "", "")
	if err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, errors.New("User not found")
	}

	user := result[0]
	userId := user["user_id"].(int64)

	// เรียก update user_id ใน patient table

	if role == "patient" {
		whereCondition = "patient_id = $1 AND user_id IS NULL"
	} else if role == "HR" || role == "medical_personnel" {
		whereCondition = "employee_id = $1 AND user_id IS NULL"
	} else {
		return 0, errors.New("Invalid role")
	}

	whereArgs = []interface{}{id}

	updateResult, err := UpdateData(table, map[string]interface{}{"user_id": userId}, whereCondition, whereArgs)
	if err != nil {
		return 0, err
	}

	if updateResult == 0 {
		return 0, errors.New("Failed to update user_id")
	}

	return updateResult, nil
}

func AuthenticateUser(username string, password string) (*auth.Token, error) {
	fields := []string{"user_id", "username", "password", "role"}
	whereCondition := "username =$1"
	whereArgs := []interface{}{username}

	userQueryResult, err := SelectData("users", fields, true, whereCondition, whereArgs, false, "", "")
	if err != nil {
		return nil, err
	}

	if len(userQueryResult) == 0 {
		return nil, errors.New("User not found")
	}

	user := userQueryResult[0]
	userId := user["user_id"].(int64)
	storedPassword := user["password"].(string)
	role := string(user["role"].([]uint8))

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return nil, errors.New("Invalid password")
	}

	if role == "patient" {
		fields := []string{"patient_id"}
		whereCondition = "user_id = $1"
		whereArgs = []interface{}{userId}
		patientQueryResult, err := SelectData("Patient", fields, true, whereCondition, whereArgs, false, "", "")
		if err != nil {
			return nil, err
		}

		patientId := patientQueryResult[0]["patient_id"].(string)

		generateJWTParam := auth.GenerateJWTClaimsParams{
			Username:  username,
			Role:      role,
			PatientID: patientId,
		}
		token, err := middlewares.GenerateJWT(generateJWTParam)
		if err != nil {
			return nil, errors.New("Failed to generate token")
		}

		return &auth.Token{
			Token: token,
		}, nil
	}

	generateJWTParam := auth.GenerateJWTClaimsParams{
		Username: username,
		Role:     role,
	}

	token, err := middlewares.GenerateJWT(generateJWTParam)
	if err != nil {
		return nil, errors.New("Failed to generate token")
	}

	return &auth.Token{
		Token: token,
	}, nil
}
