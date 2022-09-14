package controllers

import (
	"latihan/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var books []models.Book

func getIndex(id int) (int, bool) {
	for i, v := range books {
		if id == v.ID {
			return i, true
		}
	}
	return 0, false
}

func GetBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   books,
	})
}

func GetBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	index, flag := getIndex(id)
	if !flag {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  false,
			"message": "data not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   books[index],
	})
}

func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	book.ID = len(books) + 1

	books = append(books, book)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   book,
	})
}

func UpdateBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	index, flag := getIndex(id)
	if !flag {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  false,
			"message": "data not found",
		})
	}

	updateBook := models.Book{}
	c.Bind(&updateBook)
	updateBook.ID = id

	books[index] = updateBook
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   books[index],
	})
}

func DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	index, flag := getIndex(id)
	if !flag {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  false,
			"message": "data not found",
		})
	}
	books = append(books[:index], books[index+1:]...)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   books,
	})
}
