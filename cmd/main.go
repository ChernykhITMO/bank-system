package main

import (
	"bankSystem/docs"
	"bankSystem/internal/handler"
	"bankSystem/internal/model"
	"bankSystem/internal/repository"
	"bankSystem/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

func connectWithRetry(dsn string, maxAttempts int) *gorm.DB {
	var db *gorm.DB
	var err error

	for i := 1; i <= maxAttempts; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Connected to DB")
			return db
		}
		log.Printf("⏳ Attempt %d: failed to connect to DB: %v", i, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("❌ Could not connect to DB after %d attempts: %v", maxAttempts, err)
	return nil
}

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
	db := connectWithRetry(dbURL, 10)

	err := db.AutoMigrate(
		&model.UserEntity{},
		&model.AccountEntity{},
		&model.FriendsEntity{},
		&model.TransactionEntity{})

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	txManager := repository.NewGormTxManager(db)

	userRepo := repository.NewPostgresUserRepository()
	accountRepo := repository.NewPostgresAccountRepository()
	friendRepo := repository.NewPostgresFriendRepository()
	transactionRepo := repository.NewPostgresTransactionRepository()

	userService := service.NewUserService(txManager, userRepo, friendRepo)
	accountService := service.NewAccountService(txManager, accountRepo, userRepo, friendRepo, transactionRepo)

	userController := handlers.NewUserController(userService)
	accountController := handlers.NewAccountController(accountService)

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
