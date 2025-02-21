package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/models"
	"github.com/NinePTH/GO_MVC-S/src/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Register user
func Register(c echo.Context) error {
    // Enforce JSON requests
    if c.Request().Header.Get("Content-Type") != "application/json" {
        return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
    }

    // Log raw request body
    body, _ := io.ReadAll(c.Request().Body)
    fmt.Println("Raw Request Body:", string(body))
    c.Request().Body = io.NopCloser(bytes.NewBuffer(body)) // Reset body for Bind()

    var req models.AuthRequest

    // Bind request JSON to struct
    if err := c.Bind(&req); err != nil || req.Username == "" || req.Password == "" {
        fmt.Println("Bind Error:", err)
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }

    _, err := services.RegisterUser(req.Username, req.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, "User registered successfully")
}

func Login(c echo.Context) error {
    if c.Request().Header.Get("Content-Type") != "application/json" {
        return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
    }

     // Log raw request body
     body, _ := io.ReadAll(c.Request().Body)
     fmt.Println("Raw Request Body:", string(body))
     c.Request().Body = io.NopCloser(bytes.NewBuffer(body)) // Reset body for Bind()

    var req models.AuthRequest
    if err:= c.Bind(&req); err != nil || req.Username == "" || req.Password == "" {
        return c.JSON(http.StatusBadRequest, "Invalid request")
    }

    user, err := services.AuthenticateUser(req.Username, req.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, err.Error())
    }

    return c.JSON(http.StatusOK, user)
}

// Protected Profile Route
func Profile(c echo.Context) error {
    // Retrieve the "user" from the context (this is the JWT claims)
	userInterface := c.Get("user")
	claims, err := userInterface.(jwt.MapClaims)
	if !err {
		return c.JSON(http.StatusUnauthorized, "Invalid or missing user claims")
	}

	// Extract the username from the claims
	username, err := claims["username"].(string)
	if !err {
		return c.JSON(http.StatusUnauthorized, "Username not found in token claims")
	}

	// Print username for debugging (optional)
	fmt.Println("Username:", username)

    return c.JSON(http.StatusOK, map[string]string{"username":  username})
}