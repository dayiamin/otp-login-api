package utils

import (
	"fmt"
    "log"
    "math/rand"
    "time"

    "github.com/patrickmn/go-cache"
)


type OTPService struct {
    store *cache.Cache
}


// generate random number for otp with 5 numbers
func  generateOTP()(randomInt int) {
	randomInt = 10000 + rand.Intn(90000)
	
	return randomInt
}


func NewOTPService() *OTPService {
    return &OTPService{
        store: cache.New(2*time.Minute, 10*time.Minute), // TTL 5 دقیقه
    }
}

// GenerateOTP 
func (s *OTPService) GenerateOTP(phone string) string {
    otp := fmt.Sprintf("%05d", generateOTP()) // عدد 5 رقمی
    s.store.Set(phone, otp, cache.DefaultExpiration)

    log.Printf("[OTP] phone=%s otp=%s", phone, otp)
    return otp
}

// VerifyOTP takes phone number and OTP as input and returns bool
func (s *OTPService) VerifyOTP(phone, input string) bool {
    val, found := s.store.Get(phone)
    if !found || val != input {
        return false
    }
    s.store.Delete(phone)
    return true
}