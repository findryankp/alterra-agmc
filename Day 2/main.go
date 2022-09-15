package main

import (
	"latihan/config"
	"latihan/controllers"
	"latihan/middlewares"
	"latihan/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.ConnectDataBase()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	r := e.Group("/api")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &middlewares.JWTClaim{},
		SigningKey: []byte(os.Getenv("SECRETKEY")),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("user-data", controllers.UserLogin)

	routes.Routes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
