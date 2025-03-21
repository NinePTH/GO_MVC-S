package services

import (
	"fmt"

	"github.com/NinePTH/GO_MVC-S/src/models"
)



func GetUser(id string) (*models.User, error) {
	table := "users"
	fields := []string{"id", "name", "age"}

	result, err := SelectData(table, fields, true, "id = $1", []interface{}{id})

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	idRaw := string(result[0]["id"].(string))     // Convert 'id' to string
	nameRaw := string(result[0]["name"].(string)) // Convert 'name' to string
	ageRaw := int(result[0]["age"].(int64))       // Convert 'age' to int

	// Assign the values to the User struct
	var user = models.User{
		Id:   idRaw,
		Name: nameRaw,
		Age:  ageRaw,
	}

	return &user, nil
}

// GetUsers fetches a user by ID from the MySQL database
func GetAllUsers() ([]models.User, error) {
	// Define the fields
	fields := []string{"id", "name", "age"}

	// Call SelectData function
	results, err := SelectData("users", fields, false, "", nil)
	if err != nil {
		return nil, err
	}

	// Prepare a slice to hold the users
	var users []models.User

	// Iterate over the results and convert them into models.User
	for _, row := range results {

		id := string(row["id"].(string))     // Convert 'id' to string
		name := string(row["name"].(string)) // Convert 'name' to string
		age := int(row["age"].(int64))       // Convert 'age' to int

		user := models.User{
			Id:   id,
			Name: name,
			Age:  age,
		}
		users = append(users, user)
		fmt.Println(user)
	}
	fmt.Println(users)

	return users, nil
}

func AddUser(data map[string]interface{}) (int64, error) {
	table := "users"
	rowsAffected, err := InsertData(table, data)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
func UpdateUser(id string, data map[string]interface{}) (int64, error) {
	table := "users"
	condition := "id = $1"
	conditionValues := []interface{}{id}

	// Call UpdateData with correct parameters
	rowsAffected, err := UpdateData(table, data, condition, conditionValues)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}



func DeleteUser(id string) (int64, error) {
	table := "users"
	condition := "id = $1"
	conditionValues := []interface{}{id}

	rowsAffected, err := DeleteData(table, condition, conditionValues)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
