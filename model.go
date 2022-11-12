package main

import "gorm.io/gorm"

const (
	ROLE_STAFF   = "staff"
	ROLE_MANAGER = "manager"
)

type User struct {
	gorm.Model

	Username string `json:"username" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
