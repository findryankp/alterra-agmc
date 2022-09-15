package models

import "gorm.io/gorm"

var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	RoleID   string `json:"role_id" form:"role_id"`
}
