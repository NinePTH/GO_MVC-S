package controllers

import (
    "fmt"
	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/models"
	"github.com/NinePTH/GO_MVC-S/src/services"
	"github.com/labstack/echo/v4"
)

// Register user
func Register(c echo.Context) error {
    var req models.AuthRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }

    fmt.Println("req:", req)
    fmt.Printf("Username: %s, Password: %s\n",req.Username, req.Password)

    _, err := services.RegisterUser(req.Username, req.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, "User registered successfully")
}