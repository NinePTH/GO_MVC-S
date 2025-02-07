package controllers

import (
	"learn-Go/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserRoutes(e *echo.Echo) {
	e.GET("/users/:id", getUser)
	e.GET("/users", getAllUsers)
	e.POST("/users", addUser)
	e.DELETE("/users/:id", deleteUser)

	e.GET("/patient", getAllPatient)
}

func getUser(c echo.Context) error {
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

func getAllUsers(c echo.Context) error {
	user, err := services.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func addUser(c echo.Context) error {
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

func deleteUser(c echo.Context) error {
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
