package routes

import (
	"api-golang-crud/controllers"

	"api-golang-crud/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)

		// Hanya admin yang bisa menambah, mengupdate, dan menghapus user
		userGroup.Use(middleware.RoleMiddleware("admin"))
		userGroup.POST("/", controllers.CreateUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}
}
