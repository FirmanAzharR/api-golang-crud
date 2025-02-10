package routes

import (
	"api-golang-crud/controllers"
	"api-golang-crud/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/refresh", controllers.RefreshToken)
		authGroup.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)
	}
}
