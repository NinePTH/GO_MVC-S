package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/NinePTH/GO_MVC-S/src/models/auth"
)

var jwtSecret = []byte("supersecretkey")

// Generate JWT Token
func GenerateJWT(userInfo auth.GenerateJWTClaimsParams) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": userInfo.Username,
        "role": userInfo.Role,
		"patient_id": userInfo.PatientID, // If user is not patient, this will be empty
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString(jwtSecret)
}

// JWTMiddleware returns an Echo middleware function for handling JWT authentication
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
			}

			// Extract token from "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
			}

			tokenStr := parts[1]
			claims := jwt.MapClaims{}

			token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			// Store claims in context
			c.Set("user", claims)
			return next(c)
		}
	}
}