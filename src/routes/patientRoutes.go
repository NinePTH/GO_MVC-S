package routes

import (
	"github.com/NinePTH/GO_MVC-S/src/controllers"
	"github.com/NinePTH/GO_MVC-S/src/middlewares"

	"github.com/labstack/echo/v4"
)

func PatientRoutes(e *echo.Echo) {
	protected := e.Group("/patient")
	protected.Use(middlewares.JWTMiddleware()) // Apply JWT middleware

	protected.GET("", controllers.GetAllPatients) // Protected route
	protected.GET("/:id", controllers.GetPatient) // Protected route
	protected.POST("/add-patient", controllers.AddPatient) // Protected route
}