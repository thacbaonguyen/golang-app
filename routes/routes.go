package routes

import (
	"github.com/gin-gonic/gin"
	"go-ginapp/controllers"
	"go-ginapp/middleware"
)

func SetupRoutes(router *gin.Engine,
	authController controllers.AuthController,
) {

	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
		}

		//users := v1.Group("/users")
		//{
		//	users.POST("/login", authController.Login)
		//	users.POST("/register", authController.Register)
		//}
		//
		//posts := v1.Group("/posts")
		//{
		//	posts.POST("/login", authController.Login)
		//	posts.POST("/register", authController.Register)
		//}
	}
}
