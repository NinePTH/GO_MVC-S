package routes

import (
	"learn-Go/src/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
    e.GET("/users/:id", controllers.GetUser)
    e.GET("/users", controllers.GetAllUsers)
	e.GET("/patient", controllers.GetAllPatients)
	e.POST("/users",controllers. AddUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}