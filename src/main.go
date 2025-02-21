package main

import (
	"fmt"

	"github.com/NinePTH/GO_MVC-S/src/routes"
	"github.com/NinePTH/GO_MVC-S/src/utils/databaseConnector"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	databaseConnector.InitDB()

	e := echo.New()

	e.Use(middleware.CORS()) // Apply CORS for outside domain requests
    e.Use(middleware.Logger())   // Logs each request
    e.Use(middleware.Recover())  // Recovers from panics
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10))) // Limit to 10 requests per second

	// For small project we can use this way of routing, but in medium to large project we must use centralized route
	routes.UserRoutes(e)
	routes.PatientRoutes(e)
	routes.AuthRoutes(e)

	fmt.Println("Server path is http://localhost:1323/")
	e.Logger.Fatal(e.Start(":1323"))
}