package handler

import (
	"backendstartup/helper"
	"backendstartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

//userHandler akan membutuhkan bantuan (punya depedensi) dari service
type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	//map input dari user ke struct RegisterUserInput
	//struct diatas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	//Manage error validation input user
	if err != nil {
		errors := helper.FormatValidationError(err)
		//map dalam gin -> gin.H
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "tokenstatis")
	response := helper.APIResponse("Account has been register", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	//User memasukkan email dan password
	//input ditangkkap handler
	//mapping dari login ke input struct
	//input struct parsing service
	//di service cari dg bantuan repo user dan email
	//pencocokan password

	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		//map dalam gin -> gin.H
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokenmasihstatis")
	response := helper.APIResponse("Login successful", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
