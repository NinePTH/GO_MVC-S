package routes

import (
	"learn-Go/src/controllers"

	"github.com/labstack/echo/v4"
)

func PatientRoutes(e *echo.Echo) {
	e.GET("/patient", controllers.GetAllPatients)
}