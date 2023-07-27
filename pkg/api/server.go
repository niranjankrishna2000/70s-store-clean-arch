package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "main/cmd/api/docs"
	handler "main/pkg/api/handler"
	"main/pkg/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, adminHandler *handler.AdminHandler) *ServerHTTP {
	fmt.Println("=====server started=====")
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))


	routes.UserRoutes(engine.Group("/users"), userHandler, otpHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler,  userHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":1243")
}
