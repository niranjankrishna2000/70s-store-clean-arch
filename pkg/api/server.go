package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "main/cmd/api/docs"
	handler "main/pkg/api/handler"
	middleware "main/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, otpHandler *handler.OtpHandler) *ServerHTTP {
	fmt.Println("=====server started=====")
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	engine.POST("/login", userHandler.Login)
	engine.POST("/signup", userHandler.SignUp)
	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)
	// Auth middleware
	engine.Use(middleware.UserAuthMiddleware)
	{
		profile := engine.Group("/user/profile")
		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddresses)
			profile.POST("/address/add", userHandler.AddAddress)

			edit := profile.Group("/edit")
			{
				edit.PUT("/name", userHandler.EditName)
				edit.PUT("/email", userHandler.EditEmail)
				edit.PUT("/phone", userHandler.EditPhone)
			}

			security := profile.Group("/security")
			{
				security.PUT("/change-password", userHandler.ChangePassword)
			}
		}
	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":1243")
}
