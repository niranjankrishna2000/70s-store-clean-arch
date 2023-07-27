package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizationMiddleware(c *gin.Context) {
	fmt.Println("Middleware working......")
	token, _ := c.Cookie("Authorization")
	fmt.Println("Token::", token)
	fmt.Println(token)
	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
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
