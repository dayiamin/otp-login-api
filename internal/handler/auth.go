package handler

import (
	"fmt"
	"net/http"

	"github.com/dayiamin/otp-login-api/internal/database"
	"github.com/dayiamin/otp-login-api/internal/models"
	"github.com/dayiamin/otp-login-api/internal/utils"
	"github.com/gin-gonic/gin"
	"time"
)

var otpService = utils.NewOTPService()
var limiter = utils.NewRateLimiter(50 * time.Second)

// GetOTP godoc
// @Summary Request OTP
// @Description Send OTP to the user's phone (printed in logs).
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body models.RequestOTPBody true "Phone Number"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/request-otp [post]
func GetOTP(c *gin.Context) {

	var req models.RequestOTPBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if !limiter.Allow(req.PhoneNumber) {
        c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many OTP requests. Try again later."})
        return
    }


	otp := otpService.GenerateOTP(req.PhoneNumber)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("OTP sent (%s)", otp)})
	
}





// OTPLogin handles login/signup using OTP.
//
// @Summary      Verify OTP and Login or Register User
// @Description  If OTP is valid, generates JWT token and either logs in the existing user or creates a new user.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      models.OTPLogin  true  "OTP verification request"
// @Success      200      {object}  map[string]interface{}  "Login successful"
// @Success      201      {object}  map[string]interface{}  "New user created"
// @Failure      400      {object}  map[string]string        "Invalid input or OTP"
// @Failure      500      {object}  map[string]string        "Internal server error"
// @Router       /auth/verify-otp [post]
func OTPLogin(c *gin.Context) {

	var req models.OTPLogin
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	



	otpStatus := otpService.VerifyOTP(req.PhoneNumber, req.OTP)

	if !otpStatus{
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP is Invalid"})
		return
	}

	jwt ,err := utils.GenerateJWT(req.PhoneNumber)
	if err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT generating failed"})
		return
	}

	var existingProfile models.User
	if err := database.DB.Where("phone_number = ?", req.PhoneNumber).First(&existingProfile).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "logged in","JWT":jwt})
		return
	}

	newUser := models.User{
		PhoneNumber: req.PhoneNumber,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create User"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New User created","JWT":jwt})

}
