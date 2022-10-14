package handler

import (
	"crowdfunding/helper"
	"crowdfunding/payment"
	"crowdfunding/transaction"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService transaction.Service
	paymentService     payment.Service
}

func NewTransactionHandler(transactionService transaction.Service, paymentService payment.Service) *TransactionHandler {
	return &TransactionHandler{transactionService, paymentService}
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

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {

	input := transaction.CreateTransactionInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create transaction failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Createad transaction", http.StatusOK, "success", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetNotification(c *gin.Context) {
	input := transaction.TransactionNotificationInput{}
	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = h.transactionService.ProcessPayment(input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, input)
}
