package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/mapper"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/service"
	shareddto "github.com/maxcore25/proga-po-go-transport-service/internal/shared/dto"
	httphelper "github.com/maxcore25/proga-po-go-transport-service/internal/shared/http"
)

type CardHandler struct {
	service service.CardService
}

func NewCardHandler(s service.CardService) *CardHandler {
	return &CardHandler{service: s}
}

// CreateCard godoc
// @Summary Create card
// @Tags Cards
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param card body dto.CreateCardRequest true "New card"
// @Success 201 {object} dto.CardResponse
// @Failure 400 {object} shareddto.ErrorDefaultResponse
// @Failure 500 {object} shareddto.ErrorDefaultResponse
// @Router /cards [post]
func (h *CardHandler) CreateCard(c *gin.Context) {
	var req dto.CreateCardRequest

	if !httphelper.BindJSON(c, &req) {
		return
	}

	card, err := h.service.CreateCard(req)
	if err != nil {
		httphelper.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	resp := mapper.NewCardResponse(card)

	c.JSON(http.StatusCreated, resp)
}

// GetCard godoc
// @Summary Get card by ID
// @Tags Cards
// @Produce json
// @Param id path string true "Card ID (uuid)"
// @Success 200 {object} dto.CardResponse
// @Failure 400 {object} shareddto.ErrorDefaultResponse
// @Failure 404 {object} shareddto.ErrorDefaultResponse
// @Router /cards/{id} [get]
func (h *CardHandler) GetCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		httphelper.JSONError(c, http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}

	card, err := h.service.GetCard(id)
	if err != nil {
		httphelper.JSONError(c, http.StatusNotFound, errors.New("card not found"))
		return
	}

	resp := mapper.NewCardResponse(card)

	c.JSON(http.StatusOK, resp)
}

// GetAllCards godoc
// @Summary Get all cards
// @Tags Cards
// @Produce json
// @Success 200 {array} dto.CardResponse
// @Router /cards [get]
func (h *CardHandler) GetAllCards(c *gin.Context) {
	cards, err := h.service.GetAllCards()
	if err != nil {
		httphelper.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	resp := make([]*dto.CardResponse, len(cards))
	for i, card := range cards {
		resp[i] = mapper.NewCardResponse(card)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateCardByID godoc
// @Summary Update card by ID
// @Tags Cards
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Card ID (uuid)"
// @Param card body dto.UpdateCardRequest true "Card update data"
// @Success 200 {object} shareddto.MessageDefaultResponse
// @Failure 400 {object} shareddto.ErrorDefaultResponse
// @Failure 404 {object} shareddto.ErrorDefaultResponse
// @Failure 500 {object} shareddto.ErrorDefaultResponse
// @Router /cards/{id} [patch]
func (h *CardHandler) UpdateCardByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		httphelper.JSONError(c, http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}
	var updateData dto.UpdateCardRequest
	if !httphelper.BindJSON(c, &updateData) {
		return
	}
	if err := h.service.UpdateCardByID(id, updateData); err != nil {
		httphelper.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, shareddto.MessageDefaultResponse{Message: "card updated successfully"})
}

// DeleteCardByID godoc
// @Summary Delete card by ID
// @Tags Cards
// @Produce json
// @Security BearerAuth
// @Param id path string true "Card ID (uuid)"
// @Success 204 {object} nil
// @Failure 400 {object} shareddto.ErrorDefaultResponse
// @Failure 404 {object} shareddto.ErrorDefaultResponse
// @Failure 500 {object} shareddto.ErrorDefaultResponse
// @Router /cards/{id} [delete]
func (h *CardHandler) DeleteCardByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		httphelper.JSONError(c, http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}
	if err := h.service.DeleteCardByID(id); err != nil {
		httphelper.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
