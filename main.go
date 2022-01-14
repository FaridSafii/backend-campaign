package main

import (
	"backendstartup/handler"
	"backendstartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/campaign_backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("users", userHandler.RegisterUser)
	api.POST("sessions", userHandler.Login)

	router.Run(":8088")
	//menuju struct RegisterUserInput
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Test 1"
	// userInput.Email = "test1@gmail.com"
	// userInput.Occupation = "Anak IT"
	// userInput.Password = "test"

	//userService.RegisterUser(userInput)

}
