package routes

import (
	"main/pkg/api/handler"
	"main/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler) {
	engine.POST("/login", userHandler.Login)
	engine.POST("/signup", userHandler.SignUp)
	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)
	// Auth middleware
	engine.Use(middleware.UserAuthMiddleware)
	{
		{

			profile := engine.Group("/profile")
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
	}
}
