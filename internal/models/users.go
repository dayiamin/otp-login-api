package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber  string `gorm:"unique" json:"phone_number"`
}


type RequestOTPBody struct {
	PhoneNumber string `json:"phone_number" binding:"max=12,required"`
}

type OTPLogin struct {
	OTP    string `json:"otp" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"max=12,required"`
}