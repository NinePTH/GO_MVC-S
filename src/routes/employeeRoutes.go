package routes

import (
	"github.com/NinePTH/GO_MVC-S/src/controllers"
	"github.com/NinePTH/GO_MVC-S/src/middlewares"
	"github.com/labstack/echo/v4"
)

func EmployeeRoutes(e *echo.Echo) {
	protected := e.Group("/employee")
	protected.Use(middlewares.JWTMiddleware())
	protected.GET("", controllers.GetAllEmployee) // Display all employee
}
