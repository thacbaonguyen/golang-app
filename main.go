package main

import (
	"github.com/gin-gonic/gin"
	"go-ginapp/config"
	"go-ginapp/controllers"
	"go-ginapp/middleware"
	"go-ginapp/models"
	"go-ginapp/repositories"
	"go-ginapp/routes"
	"go-ginapp/services/auth"
	"go-ginapp/services/post"
	"go-ginapp/services/user"
	"go-ginapp/utils"
	"log"
)

func main() {
	config := config.LoadConfig()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Error connect to database %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Post{})
	if err != nil {
		log.Fatalf("Error cannot migrate database %v", err)
	}

	err = utils.InitializeRoles(db)
	if err != nil {
		log.Fatalf("Failed to init roles %v", err)
	}

	// repo
	userRepository := repositories.NewUserRepositoryImpl(db)
	roleRepository := repositories.NewRoleRepositoryImpl(db)
	postRepository := repositories.NewPostRepositoryImpl(db)

	jwtService := auth.NewJWTServiceImpl(config.JWTSecret, config.JWTExpiration)
	authService := auth.NewAuthServiceImpl(userRepository, roleRepository, jwtService)
	userService := user.NewUserServiceImpl(userRepository, roleRepository)
	postService := post.NewPostServiceImpl(postRepository, userRepository)

	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	postController := controllers.NewPostController(postService)

	authMiddleware := middleware.NewAuthMiddleware(jwtService, userRepository)

	router := gin.Default()

	routes.SetupRoutes(
		router,
		authController,
		userController,
		postController,
		authMiddleware,
	)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := router.Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
