package routes

import (
	"github.com/NinePTH/GO_MVC-S/src/controllers"
	"github.com/NinePTH/GO_MVC-S/src/middlewares"

	"github.com/labstack/echo/v4"
)

func PatientRoutes(e *echo.Echo) {
	protected := e.Group("/patient")
	protected.Use(middlewares.JWTMiddleware())    // Apply JWT middleware (protected route)
	protected.GET("", controllers.GetAllPatients) // Display all patient
	protected.GET("/:id", controllers.GetPatient) // Select patient by patient_id
	protected.POST("", controllers.AddPatient)    // Add patient info
	protected.PUT("", controllers.UpdatePatient)  // Update Patient info
}