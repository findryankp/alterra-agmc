package controllers

import (
	"latihan/config"
	"latihan/middlewares"
	"latihan/models"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var user models.User

	form := models.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	if err := config.Instance.Where(&form).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tokenString, err := middlewares.GenerateToken(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func UserLogin(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTClaim)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   claims,
	})
}
