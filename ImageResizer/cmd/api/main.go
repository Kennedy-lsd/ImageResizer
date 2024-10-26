package main

import (
	"github.com/Kennedy-lsd/ImageResizer/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	api := echo.New()

	// Enable CORS
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
		AllowHeaders: []string{"Content-Type"},
	}))

	api.POST("/resize", handlers.ResizeHandler)

	api.Start(":8080")
}
