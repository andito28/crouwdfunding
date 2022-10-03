package handler

import (
	"crowdfunding/helper"
	"crowdfunding/transaction"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input = transaction.GetTransactionCampaignInput{}
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.transactionService.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Error to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of campaign transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	transactions, err := h.transactionService.GetTransactionByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Error to get transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of transaction", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
