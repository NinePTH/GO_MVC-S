package routes

import (
	"github.com/NinePTH/GO_MVC-S/src/controllers"
	"github.com/NinePTH/GO_MVC-S/src/middlewares"
	"github.com/labstack/echo/v4"
)

func EmployeeRoutes(e *echo.Echo) {
	protected := e.Group("/employee")
	protected.Use(middlewares.JWTMiddleware())                   // Apply JWT middleware (protected route)
	protected.GET("", controllers.GetAllEmployee)                // Display all employee info
	protected.GET("/:id", controllers.GetEmployee)               //Display employee info by id
	protected.POST("/add-employee", controllers.AddEmployee)     //Add employee info
	protected.PUT("/update-employee", controllers.UpdateEmployee) //Update Employee info
	protected.POST("/search-employee", controllers.SearchEmployee) // Seacrh employee by id,firstname,lastname
}
