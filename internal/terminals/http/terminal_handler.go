package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	httphelper "github.com/maxcore25/proga-po-go-transport-service/internal/shared/http"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/mapper"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/service"
)

type TerminalHandler struct {
	service service.TerminalService
}

func NewTerminalHandler(s service.TerminalService) *TerminalHandler {
	return &TerminalHandler{service: s}
}

// CreateTerminal godoc
// @Summary Create terminal
// @Tags Terminals
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param terminal body dto.CreateTerminalRequest true "New terminal"
// @Success 201 {object} dto.TerminalResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /terminals [post]
func (h *TerminalHandler) CreateTerminal(c *gin.Context) {
	var req dto.CreateTerminalRequest

	if !httphelper.BindJSON(c, &req) {
		return
	}

	terminal, err := h.service.CreateTerminal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.NewTerminalResponse(terminal)

	c.JSON(http.StatusCreated, resp)
}

// GetTerminal godoc
// @Summary Get terminal by ID
// @Tags Terminals
// @Produce json
// @Param id path string true "Terminal ID (uuid)"
// @Success 200 {object} dto.TerminalResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /terminals/{id} [get]
func (h *TerminalHandler) GetTerminal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	terminal, err := h.service.GetTerminal(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "terminal not found"})
		return
	}

	resp := mapper.NewTerminalResponse(terminal)

	c.JSON(http.StatusOK, resp)
}

// GetAllTerminals godoc
// @Summary Get all terminals
// @Tags Terminals
// @Produce json
// @Success 200 {array} dto.TerminalResponse
// @Router /terminals [get]
func (h *TerminalHandler) GetAllTerminals(c *gin.Context) {
	terminals, err := h.service.GetAllTerminals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]*dto.TerminalResponse, len(terminals))
	for i, terminal := range terminals {
		resp[i] = mapper.NewTerminalResponse(terminal)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateTerminalByID godoc
// @Summary Update terminal by ID
// @Tags Terminals
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Terminal ID (uuid)"
// @Param terminal body dto.UpdateTerminalRequest true "Terminal update data"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /terminals/{id} [patch]
func (h *TerminalHandler) UpdateTerminalByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	var updateData dto.UpdateTerminalRequest
	if !httphelper.BindJSON(c, &updateData) {
		return
	}
	if err := h.service.UpdateTerminalByID(id, updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "terminal updated successfully"})
}

// DeleteTerminalByID godoc
// @Summary Delete terminal by ID
// @Tags Terminals
// @Produce json
// @Security BearerAuth
// @Param id path string true "Terminal ID (uuid)"
// @Success 204 {object} nil
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /terminals/{id} [delete]
func (h *TerminalHandler) DeleteTerminalByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	if err := h.service.DeleteTerminalByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
