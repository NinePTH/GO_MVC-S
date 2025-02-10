package main

import (
    "fmt"
    "learn-Go/src/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"learn-Go/src/database"
)

func main() {
	database.InitDB()

	e := echo.New()

	e.Use(middleware.CORS())
	
	// For small project we can use this way of routing, but in medium to large project we must use centralized route
	routes.UserRoutes(e)
	routes.PatientRoutes(e)

	fmt.Println("Server path is http://localhost:1323/")
	e.Logger.Fatal(e.Start(":1323"))
}