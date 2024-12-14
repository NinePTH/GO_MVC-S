package main

import (
    "fmt"
    "learn-Go/src/controllers"
	"github.com/labstack/echo/v4"
	"learn-Go/src/database"
)

func main() {
	database.InitDB()

	e := echo.New()

	controllers.GetUserRoutes(e)

	fmt.Println("Server path is http://localhost:1323/")
	e.Logger.Fatal(e.Start(":1323"))
}