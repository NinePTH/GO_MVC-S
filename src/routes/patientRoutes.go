package routes

import (
	"github.com/NinePTH/GO_MVC-S/src/controllers"

	"github.com/labstack/echo/v4"
)

func PatientRoutes(e *echo.Echo) {
	e.GET("/patient", controllers.GetAllPatients)
}