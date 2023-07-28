package routes

import (
	"main/pkg/api/handler"
	"main/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, userHandler *handler.UserHandler,categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler )

	engine.Use(middleware.AdminAuthMiddleware)
	{
		usermanagement := engine.Group("/users")
		{
			usermanagement.POST("/block", adminHandler.BlockUser)
			usermanagement.POST("/unblock", adminHandler.UnBlockUser)
			usermanagement.GET("/getusers", adminHandler.GetUsers)
		}

		categorymanagement := engine.Group("/category")
		{
			categorymanagement.POST("/add", categoryHandler.AddCategory)
			categorymanagement.PUT("/update", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("/delete", categoryHandler.DeleteCategory)
		}

		inventorymanagement := engine.Group("/inventories")
		{
			inventorymanagement.POST("/add", inventoryHandler.AddInventory)
			inventorymanagement.PUT("/update", inventoryHandler.UpdateInventory)
			inventorymanagement.DELETE("/delete", inventoryHandler.DeleteInventory)
		}
		orders := engine.Group("/orders")
		{
			orders.PUT("/edit/status", orderHandler.EditOrderStatus)
			orders.GET("", orderHandler.AdminOrders)
		}
	}
}