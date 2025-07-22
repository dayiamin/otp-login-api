package main

import (
	"os"

	"github.com/dayiamin/otp-login-api/internal/database"
	"github.com/dayiamin/otp-login-api/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/dayiamin/otp-login-api/docs"
)


func init(){
	var envErr = godotenv.Load()
	if envErr != nil{
		os.Setenv("JWT_SECRET","123321aawedaf")
	}
}

// @title OTP Auth API
// @version 1.0
// @description API for OTP-based authentication.
// @host localhost:8080
// @BasePath /
func main() {

	database.Connect()
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth routes
	authGroup := r.Group("/auth")
	routes.AuthRoutes(authGroup)

	r.Run(":8080")
}
