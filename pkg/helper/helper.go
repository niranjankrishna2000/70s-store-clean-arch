package helper

import (
	"errors"
	"fmt"
	"main/pkg/utils/models"

	"os"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient

func GenerateTokenUser(user models.UserResponse) (string, error) {
	fmt.Println("---CreateToken Function Called")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Username,
		"role": "user",
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("KEY")))

	if err == nil {
		fmt.Println("token created")
	}
	return tokenString, nil
}

func PasswordHashing(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}
func TwilioSetup(username string, password string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})

}

func TwilioSendOTP(phone string, serviceID string) (string, error) {
	to := "+91" + phone
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceID, params)
	if err != nil {

		return " ", err
	}

	return *resp.Sid, nil

}

func  TwilioVerifyOTP(serviceID string, code string, phone string) error {

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)

	if err != nil {
		return err
	}

	if *resp.Status == "approved" {
		return nil
	}

	return errors.New("failed to validate otp")

}