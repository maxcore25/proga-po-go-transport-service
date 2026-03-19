package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	httphelper "github.com/maxcore25/proga-po-go-transport-service/internal/shared/http"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/mapper"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

// CreateTransaction godoc
// @Summary Create transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param transaction body dto.CreateTransactionRequest true "New transaction"
// @Success 201 {object} dto.TransactionResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req dto.CreateTransactionRequest

	if !httphelper.BindJSON(c, &req) {
		return
	}

	transaction, err := h.service.CreateTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.NewTransactionResponse(transaction)

	c.JSON(http.StatusCreated, resp)
}

// GetTransaction godoc
// @Summary Get transaction by ID
// @Tags Transactions
// @Produce json
// @Param id path string true "Transaction ID (uuid)"
// @Success 200 {object} dto.TransactionResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	transaction, err := h.service.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	resp := mapper.NewTransactionResponse(transaction)

	c.JSON(http.StatusOK, resp)
}

// GetAllTransactions godoc
// @Summary Get all transactions
// @Tags Transactions
// @Produce json
// @Success 200 {array} dto.TransactionResponse
// @Router /transactions [get]
func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]*dto.TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		resp[i] = mapper.NewTransactionResponse(transaction)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateTransactionByID godoc
// @Summary Update transaction by ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Transaction ID (uuid)"
// @Param transaction body dto.UpdateTransactionRequest true "Transaction update data"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transactions/{id} [patch]
func (h *TransactionHandler) UpdateTransactionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	var updateData dto.UpdateTransactionRequest
	if !httphelper.BindJSON(c, &updateData) {
		return
	}
	if err := h.service.UpdateTransactionByID(id, updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "transaction updated successfully"})
}

// DeleteTransactionByID godoc
// @Summary Delete transaction by ID
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Param id path string true "Transaction ID (uuid)"
// @Success 204 {object} nil
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transactions/{id} [delete]
func (h *TransactionHandler) DeleteTransactionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	if err := h.service.DeleteTransactionByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
