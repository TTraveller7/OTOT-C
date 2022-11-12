package main

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func createTable() {
	db.AutoMigrate(&User{})
}

func saveUser(u *User) error {
	return db.Save(u).Error
}

func getUserByUsername(n string) (*User, error) {
	u := &User{}
	if err := db.Model(&User{}).Where("username = ?", n).First(u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func initDb() {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: getPostgresUrl(),
	}), &gorm.Config{})
	if err != nil {
		log.Error("", err)
		os.Exit(1)
	}
	db = database
	createTable()
}
