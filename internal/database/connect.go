package database

import (
	"log"

	"github.com/dayiamin/otp-login-api/internal/models"

	"github.com/glebarez/sqlite"  //for using pure go for sql
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect(){


	db, err := gorm.Open(sqlite.Open("otp-api.db"), &gorm.Config{})
	if err != nil{
		log.Fatal("db connections failed")
	}
	db.AutoMigrate(&models.User{})
	DB = db

	
}