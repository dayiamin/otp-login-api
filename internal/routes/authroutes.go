package routes

import (
	"github.com/dayiamin/otp-login-api/internal/handler"
	"github.com/gin-gonic/gin"
)



func AuthRoutes(r *gin.RouterGroup){
	userGroup := r.Group("/v1")
	userGroup.POST("/get-otp",handler.GetOTP)
	userGroup.POST("/verify-otp",handler.OTPLogin)
	

}