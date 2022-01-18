package handler

import (
	"backendstartup/helper"
	"backendstartup/transaction"
	"backendstartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

//List transaksi
//List transaksi by User
//Create transaksi

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get transaction of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get transaction of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get campaign transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
