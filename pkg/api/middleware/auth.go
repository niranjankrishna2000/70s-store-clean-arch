package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// func AdminAuthMiddleware(c *gin.Context) {

// 	accessToken := c.Request.Header.Get("Authorization")

// 	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("accesssecret"), nil
// 	})

// 	if err == nil {
// 		c.Next()
// 	}

// 	refreshToken := c.Request.Header.Get("RefreshToken")

// 	// Check if the refresh token is valid.
// 	_, err = jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("refreshsecret"), nil
// 	})
// 	if err != nil {
// 		// The refresh token is invalid.
// 		c.AbortWithStatus(401)
// 		return
// 	}
// 	// The access token is invalid. Check the refresh token.

// 	// The refresh token is valid. Generate a new access token.
// 	newAccessToken, err := CreateNewAccessTokenAdmin()
// 	if err != nil {
// 		// An error occurred while generating the new access token.
// 		c.AbortWithStatus(500)
// 		return
// 	}

//		// Set the new access token in the response header.
//		c.Header("Authorization", "jwt "+newAccessToken)
//		c.Next()
//	}
func AdminAuthMiddleware(c *gin.Context) {
	fmt.Println("Middleware working......")
	token, _ := c.Cookie("Authorization")
	fmt.Println("Token::", token)
	fmt.Println(token)
	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}

func validateToken(token string) error {
	fmt.Println("Token validating.........")
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		secret := os.Getenv("KEY")
		return []byte(secret), nil
	})

	return err
}
