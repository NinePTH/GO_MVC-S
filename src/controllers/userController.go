package controllers

import (
    "learn-Go/src/services"
	"net/http"
    "github.com/labstack/echo/v4"
)

func GetUserRoutes(e *echo.Echo) {
    e.GET("/users/:id", getUser)
    e.GET("/users", getAllUsers)
}

func getUser(c echo.Context) error {
	id := c.Param("id")	
	user, err := services.GetUser(id)
	if err != nil {
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