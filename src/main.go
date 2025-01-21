package main

import (
    "fmt"
    "learn-Go/src/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"learn-Go/src/database"
)

func main() {
	database.InitDB()

	e := echo.New()

	e.Use(middleware.CORS())
	
	controllers.GetUserRoutes(e)

	fmt.Println("Server path is http://localhost:1323/")
	e.Logger.Fatal(e.Start(":1323"))
}