package main

import (
	"backendstartup/auth"
	"backendstartup/campaign"
	"backendstartup/handler"
	"backendstartup/helper"
	"backendstartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/campaign_backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//Auth
	authService := auth.NewService()

	//User
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	//Campaign
	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	//fmt.Println(authService.GenerateToken(1001))
	router := gin.Default()
	//Routing get images file on client side, params 1 routenya, params 2 lokasi folder
	router.Static("/images", "./images")

	api := router.Group("/api/v1")
	//userService.SaveAvatar(1, "images/1-profile.png")
	api.POST("users", userHandler.RegisterUser)
	api.POST("sessions", userHandler.Login)
	api.POST("email_checker", userHandler.CheckEmailAvailability)
	api.GET("campaigns", campaignHandler.GetCampaigns)
	api.POST("campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	//endpoint get campaign by ID
	api.GET("campaigns/:id", campaignHandler.GetCampaign)

	//Update Campaign
	api.PUT("campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)

	//endpoint upload avatars
	api.POST("avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run(":8088")

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Mengambil header
		authHeader := c.GetHeader("Authorization")
		//jika header tidak memiliki string Bearer maka Unauthorized
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		//memisahkan string dengan karakter spasi
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//Mencek apakah method yang digunakan sama menggunakan ValidateToken
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//Jika metode sama maka dicek apakah token Valid atau tidak
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Jika sudah melewati validasi diatas maka mengambil data user_id di claim jwt
		userID := int(claim["user_id"].(float64))
		//dan dicari kedalam service GetUserByID
		user, err := userService.GetUserByID(userID)
		//Jika tidak ada data user maka Unauthorized
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//data yang ditemukan di claim akan di set di GIN kedalam variabel currentUser -> handler
		c.Set("currentUser", user)
	}
}
