package main

import (
	"github.com/gin-gonic/gin"
	"go-ginapp/config"
	"go-ginapp/models"
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
	//userRepository := repositories.NewUserRepositoryImpl(db)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
