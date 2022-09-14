package models

type Book struct {
	ID     int    `json:"id" form:"id"`
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Year   int    `json:"year" form:"year"`
}
