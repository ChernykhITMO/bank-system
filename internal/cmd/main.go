package main

import (
	"bankSystem/internal/domain/enums"
	"bankSystem/internal/model"
	"bankSystem/internal/repostitory"
	"bankSystem/internal/service"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

func main() {
	// 1. Загрузка .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Подключение к базе
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// 3. Миграция таблиц
	err = db.AutoMigrate(&model.UserEntity{}, &model.AccountEntity{}, &model.TransactionEntity{}, &model.FriendsEntity{})
	if err != nil {
		log.Fatal("failed to migrate:", err)
	}

	// 4. Репозитории и сервисы
	userRepo := repostitory.NewPostgresUserRepository(db)
	accountRepo := repostitory.NewPostgresAccountRepository(db)
	friendRepo := repostitory.NewPostgresFriendRepository(db)

	userService := service.NewUserService(userRepo, friendRepo)
	accountService := service.NewAccountService(accountRepo, userRepo, friendRepo)

	// 5. CLI интерфейс
	for {
		fmt.Println("\n=== BANK SYSTEM ===")
		fmt.Println("1. Create user")
		fmt.Println("2. Add friend")
		fmt.Println("3. Create account")
		fmt.Println("4. Deposit")
		fmt.Println("5. Withdraw")
		fmt.Println("6. Transfer")
		fmt.Println("7. Show balance")
		fmt.Println("8. Show history")
		fmt.Println("0. Exit")
		fmt.Print("Choose option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var login, name, sexInput, colorInput string
			fmt.Print("Login: ")
			fmt.Scan(&login)
			fmt.Print("Name: ")
			fmt.Scan(&name)
			fmt.Print("Sex (male/female): ")
			fmt.Scan(&sexInput)
			fmt.Print("Hair color (black/white): ")
			fmt.Scan(&colorInput)

			sex := enums.Sex(strings.ToLower(sexInput))
			color := enums.Color(strings.ToLower(colorInput))

			_, err := userService.NewUser(login, name, sex, color)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("User created!")
			}

		case 2:
			var login1, login2 string
			fmt.Print("User login: ")
			fmt.Scan(&login1)
			fmt.Print("Friend login: ")
			fmt.Scan(&login2)

			user1, err1 := userService.GetUser(login1)
			user2, err2 := userService.GetUser(login2)

			if err1 == nil && err2 == nil {
				userService.AddFriend(user1.Login, user2.Login)
				fmt.Println("Friends added.")
			} else {
				fmt.Println("User(s) not found.")
			}

		case 3:
			var login string
			fmt.Print("Login: ")
			fmt.Scan(&login)

			user, err := userService.GetUser(login)
			if err == nil {
				accountService.NewUserAccount(user)
				fmt.Println("Account created.")
			} else {
				fmt.Println("User not found.")
			}
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
