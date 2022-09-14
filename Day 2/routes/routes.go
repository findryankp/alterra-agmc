package routes

import (
	"os"

	"github.com/labstack/echo/v4"
)

func Routes() {
	e := echo.New()
	e.Start(os.Getenv("PORT"))
}
