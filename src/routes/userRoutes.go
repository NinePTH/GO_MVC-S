package routes

import (
	"github.com/NinePTH/GO_MVC-S/src/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	e.GET("/users/:id", controllers.GetUser)
	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.AddUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
