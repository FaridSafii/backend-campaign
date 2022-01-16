package handler

import (
	"backendstartup/campaign"
	"backendstartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// note analisis
// tangkap parameter di handler
// handler set data ke service
// service yang menentukan repo mana yang digunakan
// repo access db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	//Query parameter dengan menangkap params dari get
	//hasil string di konversi ke integer menggunakan strconv.Atoi
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Successfully to get campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}
