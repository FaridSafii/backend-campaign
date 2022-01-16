package main

import (
	"backendstartup/auth"
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
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	//fmt.Println(authService.GenerateToken(1001))
	router := gin.Default()
	api := router.Group("/api/v1")
	//userService.SaveAvatar(1, "images/1-profile.png")
	api.POST("users", userHandler.RegisterUser)
	api.POST("sessions", userHandler.Login)
	api.POST("email_checker", userHandler.CheckEmailAvailability)
	api.POST("avatars", userHandler.UploadAvatar)
	router.Run(":8088")
	//menuju struct RegisterUserInput
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Test 1"
	// userInput.Email = "test1@gmail.com"
	// userInput.Occupation = "Anak IT"
	// userInput.Password = "test"

	//userService.RegisterUser(userInput)

}
