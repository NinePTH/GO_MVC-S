package controllers

import (

	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/services"

	"github.com/labstack/echo/v4"
)





func UpdateUser(c echo.Context) error {
	id := c.QueryParam("id") // Get the user ID from the query parameter
	if id == "" {
		return c.JSON(http.StatusBadRequest, "Missing user ID")
	}

	name := c.QueryParam("name") // Get the name from the query parameter
	age := c.QueryParam("age")   // Get the age from the query parameter
	//ดัก null
	if name == "" {
		return c.JSON(http.StatusBadRequest, "Missing User Name")
	}
	if age == "" {
		return c.JSON(http.StatusBadRequest, "Missing Age")
	}

	// Create the data map for the update
	data := map[string]interface{}{
		"name": name,
		"age":  age,
	}

	rowsAffected, err := services.UpdateUser(id, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, "No rows affected")
	}

	return c.JSON(http.StatusOK, "User updated successfully")
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := services.GetUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNoContent, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func GetAllUsers(c echo.Context) error {
	user, err := services.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	age := c.QueryParam("age")

	data := map[string]interface{}{"id": id, "name": name, "age": age}

	rowsAffected, err := services.AddUser(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, "No rows affected")
	}

	c.JSON(http.StatusOK, "User added successfully")

	return nil
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	rowsAffected, err := services.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, "No rows affected")
	}

	c.JSON(http.StatusOK, "User deleted successfully")

	return nil
}
