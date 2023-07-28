package routes

import (
	"main/pkg/api/handler"
	"main/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, inventoryHandler *handler.InventoryHandler, cartHandler *handler.CartHandler,orderHandler *handler.OrderHandler) {
	engine.POST("/login", userHandler.Login)
	engine.POST("/signup", userHandler.SignUp)
	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)
	// Auth middleware
	engine.Use(middleware.UserAuthMiddleware)
	{

		search := engine.Group("/search")
		{
			search.POST("/", inventoryHandler.SearchProducts)
		}

		home := engine.Group("/home")
		{
			home.GET("/products", inventoryHandler.ListProducts)
			home.GET("/products/details", inventoryHandler.ShowIndividualProducts)
			home.POST("/add-to-cart", cartHandler.AddToCart)

		}
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
			orders := profile.Group("/orders")
			{
				orders.GET("", orderHandler.GetOrders)
				orders.DELETE("", orderHandler.CancelOrder)
				
			}
		}
		
		cart := engine.Group("/cart")
		{
			cart.GET("/", userHandler.GetCart)
			cart.DELETE("/remove", userHandler.RemoveFromCart)
			cart.PUT("/updateQuantity/plus", userHandler.UpdateQuantityAdd)
			cart.PUT("/updateQuantity/minus", userHandler.UpdateQuantityLess)

		}
		
		checkout := engine.Group("/check-out")
		{
			checkout.GET("", cartHandler.CheckOut)
			checkout.POST("/order", orderHandler.OrderItemsFromCart)
		}

	}
}
