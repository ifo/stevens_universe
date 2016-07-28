package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name string
	Pass string
	Salt string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "Steve", Pass: "TODO hash and salt"})

	// Read
	var user User
	db.First(&user, 1)                   // find user with id 1
	db.First(&user, "name = ?", "Steve") // find user with name Steve

	// Update - update user's pass
	db.Model(&user).Update("Pass", "Seriously hash and salt this")

	// Delete - delete user
	//db.Delete(&user)
}
