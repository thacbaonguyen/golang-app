package routes

import (
	"github.com/gin-gonic/gin"
	"go-ginapp/controllers"
	"go-ginapp/middleware"
)

func SetupRoutes(router *gin.Engine,
	authController controllers.AuthController,
	userController controllers.UserController,
	postController controllers.PostController,
	authMiddleware middleware.AuthMiddleware,
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

		users := v1.Group("/users")
		{
			users.GET("/:id", userController.GetUserByID)
			users.Use(authMiddleware.AuthRequired())
			users.GET("/all", userController.GetAllUsers)
			users.GET("/me", userController.GetCurrentUser)
			users.PUT("/update", userController.UpdateUser)
			users.PUT("/update/password", userController.ChangePassword)
			users.DELETE("/delete/:userId", userController.DeleteUser)
		}

		posts := v1.Group("/posts")
		{
			// public
			posts.GET("/all", postController.GetAllPosts)
			posts.GET("/:postId", postController.GetPostByID)
			posts.GET("/user/:userId", postController.GetPostsByUser)
			//protect
			posts.Use(authMiddleware.AdminRequired())
			posts.POST("/create", postController.CreatePost)
			posts.PUT("/update/:postId", postController.UpdatePost)
			posts.DELETE("/delete/:postId", postController.DeletePost)
		}
	}
}
