package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"main/pkg/utils/models"
)

// type userKey string
// func (k userKey) String() string {
// 	return string(k)
// }

/*
UserAuthMiddleware is a middleware for user authentication

Parameters:
- c: Gin Context.
*/
func UserAuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := viper.GetString("KEY")
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	fmt.Println("==========", claims)
	role, ok := claims["role"].(string)
	if !ok || role != "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	id, ok := claims["userid"].(float64)
	fmt.Println("============", id)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access id"})
		c.Abort()
		return
	}
	fmt.Println("============", id)
	userIDString := fmt.Sprintf("%v", id)
	//c.SetCookie("userID", userIDString, 3600, "", "", true, true)
	var key models.UserKey = "userID"
	var val models.UserKey = models.UserKey(userIDString)

	fmt.Println("userKey is ", key)
	fmt.Println("value is ", val)

	ctx := context.WithValue(c, key, val)
	fmt.Println("ctx is", ctx)
	// Set the context to the request
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
