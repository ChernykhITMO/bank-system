package main

import (
	"bankSystem/controller"
	"bankSystem/docs"
	repostitory2 "bankSystem/repostitory"
	"bankSystem/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

// @title Bank SystemAPI
// @version 1.0
// @description This is a simple banking API for practice
// @host localhost:8080
// // @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env", err)
	}

	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	userRepo := repostitory2.NewPostgresUserRepository(db)
	friendRepo := repostitory2.NewPostgresFriendRepository(db)
	userService := service.NewUserService(userRepo, friendRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, it's bank system!"})
	})

	router.POST("/user/create", userController.CreateUser)
	router.POST("user/add_friend", userController.AddFriend)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
