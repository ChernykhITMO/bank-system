package main

import (
	"bankSystem/controller"
	"bankSystem/docs"
	"bankSystem/model"
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

// @title Bank System API
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
	db.AutoMigrate(&model.UserEntity{}, &model.AccountEntity{}, &model.FriendsEntity{}, &model.TransactionEntity{})

	userRepo := repostitory2.NewPostgresUserRepository(db)
	accountRepo := repostitory2.NewPostgresAccountRepository(db)
	friendRepo := repostitory2.NewPostgresFriendRepository(db)
	transactionRepo := repostitory2.NewPostgresTransactionRepository(db)
	userService := service.NewUserService(userRepo, friendRepo)
	accountService := service.NewAccountService(accountRepo, userRepo, friendRepo, transactionRepo)
	userController := controller.NewUserController(userService)
	accountController := controller.NewAccountController(accountService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, it's bank system!"})
	})

	router.POST("/user/create", userController.CreateUser)
	router.POST("/user/add_friend", userController.AddFriend)
	router.POST("/user/remove_friend", userController.RemoveFriend)
	router.GET("/user/get_user", userController.GetUser)
	router.POST("/account/create", accountController.CreateAccount)
	router.GET("/account/balance", accountController.GetBalance)
	router.POST("/account/deposit", accountController.Deposit)
	router.POST("/account/withdraw", accountController.Withdraw)
	router.POST("/account/transfer", accountController.Transfer)
	router.DELETE("/account/delete", accountController.DeleteAccount)
	router.GET("/account/transactions", accountController.GetTransactions)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
