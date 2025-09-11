package services

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type OTPService struct{}

func NewOTPService() *OTPService {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	return &OTPService{}
}

// GenerateOTP generates a 6-digit OTP code
func (s *OTPService) GenerateOTP() string {
	// Generate random 6-digit number
	otp := rand.Intn(999999-100000+1) + 100000
	return fmt.Sprintf("%06d", otp)
}

// SendSMS simulates sending SMS with OTP
// In production, you would integrate with SMS providers like Twilio, AWS SNS, etc.
func (s *OTPService) SendSMS(phone, otp string) error {
	// Simulate SMS sending delay
	time.Sleep(100 * time.Millisecond)
	
	// Log the OTP for development/testing purposes
	log.Printf("📱 SMS Sent to %s: Your OTP code is %s (Valid for 5 minutes)", phone, otp)
	
	// In production, you would implement actual SMS sending here:
	// Example with Twilio:
	// client := twilio.NewClient(accountSid, authToken)
	// message := &twilio.CreateMessageParams{
	//     From: "+1234567890",
	//     To:   phone,
	//     Body: fmt.Sprintf("Your verification code is: %s", otp),
	// }
	// _, err := client.CreateMessage(message)
	// return err
	
	// For now, simulate successful sending
	return nil
}

// ValidateOTPExpiry checks if OTP is still valid (within 5 minutes)
func (s *OTPService) ValidateOTPExpiry(createdAt time.Time) bool {
	expiryDuration := 5 * time.Minute
	return time.Since(createdAt) <= expiryDuration
}

// CleanupExpiredOTPs removes expired OTP records (should be called periodically)
func (s *OTPService) GetExpiryTime() time.Time {
	return time.Now().Add(5 * time.Minute)
}