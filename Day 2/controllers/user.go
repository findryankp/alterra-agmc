package controllers

import (
	"latihan/config"
	"latihan/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func checkUserExist(user *models.User, id int) (bool, string) {
	if err := config.Instance.Find(&user, id).Error; err != nil {
		return false, err.Error()
	}
	return true, ""
}

func GetUsers(c echo.Context) error {
	var users []models.User

	if err := config.Instance.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   users,
	})
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	data, err := checkUserExist(&user, id)
	if !data {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   user,
	})
}

func CreateUser(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	config.Instance.Create(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   user,
	})
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	data, err := checkUserExist(&user, id)
	if !data {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	formUpdate := c.Bind(&user)
	config.Instance.Model(&user).Updates(formUpdate)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   user,
	})
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	data, err := checkUserExist(&user, id)
	if !data {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	config.Instance.Delete(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   "",
	})
}
