package helper

import (
	"context"
	"errors"
	"fmt"
	"log"
	"main/pkg/domain"
	"main/pkg/utils/models"
	"mime/multipart"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
	"golang.org/x/crypto/bcrypt"
)

var client *twilio.RestClient

/*
GenerateTokenUser creates a jwt token for user

Parameters:
- user: user details

Returns:
- string: JWT token string
- error: error is returned
*/
func GenerateTokenUser(user models.UserResponse) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":   user.Username,
		"role":   "user",
		"userid": user.Id,
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("KEY")))

	if err == nil {
		fmt.Println("token created")
	}
	return tokenString, nil
}

/*
PasswordHashing hashes a password.

Parameters:
- password: Password to be hashed.

Returns:
- string: Hashed Password.
- error: Error is returned if any.
*/
func PasswordHashing(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

/*
TwilioSetup will setup the twillio.

Parameters:
- username: Twillio Username.
- password: Twillio Password.
*/
func TwilioSetup(username string, password string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})

}

/*
TwilioSendOTP sends otp to the number provides from the specified service

Parameters:
- phone: Otp reciever phone number.
- serviceID: Twillio Service ID to choose the service.

Returns:
- string: The unique string that we created to identify the Verification resource.
- error: Error is returned if any.
*/
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

/*
TwilioVerifyOTP verifies the otp sent to the number

Parameters:
- phone: Otp reciever phone number.
- serviceID: Twillio Service ID to choose the service.
- code: OTP.

Returns:
- error: Error is returned if any.
*/
func TwilioVerifyOTP(serviceID string, code string, phone string) error {

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

/*
GenerateTokenAdmin creates a jwt token for admin

Parameters:
- admin: admin details

Returns:
- string: JWT token string
- error: error is returned
*/
func GenerateTokenAdmin(admin domain.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": admin.Username,
		"role":  "admin",
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("KEY")))

	if err == nil {
		fmt.Println("token created")
	}
	return tokenString, nil
}

// func GetUserID(c *gin.Context) (int, error) {
// 	var key models.UserKey = "userID"
// 	var val models.UserKey = c.Request.Context().Value(key).(models.UserKey)

// 	ID := val.String()
// 	userID, _ := strconv.Atoi(ID)
// 	return userID,nil
// }

/*
GetUserID returns the userID stored in the context

Parameters:
- c: gin context

Returns:
- int: userID
- error: error is returned
*/
func GetUserID(c *gin.Context) (int, error) {
	var key models.UserKey = "userID"
	val := c.Request.Context().Value(key)

	// Check if the value is not nil
	if val == nil {
		return 0, errors.New("userID not found in context")
	}

	// Use type assertion to convert to the expected type
	userKey, ok := val.(models.UserKey)
	if !ok {
		return 0, errors.New("failed to convert userID to the expected type")
	}

	ID := userKey.String()
	userID, err := strconv.Atoi(ID)
	if err != nil {
		return 0, errors.New("failed to convert userID to int")
	}

	return userID, nil
}

func FindMostBoughtProduct(products []domain.ProductReport) []int {

	productMap := make(map[int]int)

	for _, v := range products {
		productMap[v.InventoryID] += v.Quantity
	}

	maxQty := 0
	for _, v := range productMap {
		if v > maxQty {
			maxQty = v
		}
	}

	var bestSellers []int
	for k, v := range productMap {
		if v == maxQty {
			bestSellers = append(bestSellers, k)
		}
	}
	return bestSellers
}

func AddImageToS3(file *multipart.FileHeader) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}
	// creds := credentials.NewStaticCredentialsProvider("AKIATXR23JPSQ3U2OH5G", "AKIATXR23JPSQ3U2OH5G", "your-session-token")

	cfg, err := config.LoadDefaultConfig(context.TODO() /*config.WithCredentialsProvider(creds), */, config.WithRegion("ap-southeast-2"))
	if err != nil {
		fmt.Println("configuration error:", err)
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	f, openErr := file.Open()
	if openErr != nil {
		fmt.Println("opening error:", openErr)
		return "", openErr
	}
	defer f.Close()

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("seventysstore"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if uploadErr != nil {
		fmt.Println("uploading error:", uploadErr)
		return "", uploadErr
	}

	return result.Location, nil
}
