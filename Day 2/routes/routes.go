package routes

import (
	"latihan/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/books", controllers.GetBooks)
	e.GET("/books/:id", controllers.GetBook)
	e.POST("/books", controllers.CreateBook)
	e.PUT("/books/:id", controllers.UpdateBook)
	e.DELETE("/books/:id", controllers.DeleteBook)

	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	e.POST("/login", controllers.Login)
}
