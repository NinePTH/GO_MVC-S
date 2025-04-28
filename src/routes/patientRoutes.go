package routes

import (
	"github.com/NinePTH/GO_MVC-S/src/controllers"
	"github.com/NinePTH/GO_MVC-S/src/middlewares"

	"github.com/labstack/echo/v4"
)

func PatientRoutes(e *echo.Echo) {
	protected := e.Group("/patient")
	protected.Use(middlewares.JWTMiddleware())    // Apply JWT middleware (protected route)
	protected.GET("", controllers.GetAllPatients) // Display all patient info
	protected.GET("/:id", controllers.GetPatient) // Select patient info by patient_id
	protected.PUT("/update-patient", controllers.UpdatePatient)  // Update Patient info
	protected.POST("/add-patient", controllers.AddPatient) // Add patient info
	protected.POST("/add-patient-history", controllers.AddPatientHistory) // Add patient history
	protected.POST("/add-patient-appointment", controllers.AddPatientAppointment) // Add patient appointment
}