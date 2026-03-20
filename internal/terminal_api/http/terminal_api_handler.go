package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/maxcore25/proga-po-go-transport-service/internal/shared/dto"
	httphelper "github.com/maxcore25/proga-po-go-transport-service/internal/shared/http"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminal_api/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminal_api/service"
)

type TerminalAPIHandler struct {
	service service.TerminalAPIService
}

func NewTerminalAPIHandler(s service.TerminalAPIService) *TerminalAPIHandler {
	return &TerminalAPIHandler{service: s}
}

// AuthorizePayment godoc
// @Summary      Authorize payment transaction
// @Description  The terminal sends the card UID and amount. The server runs all
// @Description  business checks in order and returns an approved/declined result.
// @Description  Declined transactions are also persisted for audit purposes.
// @Tags         Terminal API
// @Accept       json
// @Produce      json
// @Param        body body dto.PaymentAuthRequest true "Payment authorization request"
// @Success      200 {object} dto.PaymentAuthResponse "Approved or declined — HTTP 200 in both cases; check 'approved' field"
// @Failure      400 {object} shareddto.ErrorDefaultResponse                   "Validation error (missing fields, amount ≤ 0, etc.)"
// @Failure      500 {object} shareddto.ErrorDefaultResponse                   "Internal server error (DB write failure)"
// @Router       /terminal/authorize [post]
func (h *TerminalAPIHandler) AuthorizePayment(c *gin.Context) {
	var req dto.PaymentAuthRequest
	if !httphelper.BindJSON(c, &req) {
		return
	}

	// Бизнес-отказ (card_blocked, insufficient_funds и т.д.) — всегда HTTP 200.
	// Только настоящие серверные ошибки дают 500.
	resp, err := h.service.AuthorizePayment(req)
	if err != nil {
		httphelper.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// LoadKeys godoc
// @Summary      Load encryption keys for terminal
// @Description  Returns all active MIFARE encryption keys. The terminal calls this
// @Description  on startup and caches the keys locally to decrypt card sectors.
// @Description  Only active terminals (identified by serial_number) can load keys.
// @Tags         Terminal API
// @Produce      json
// @Param        terminal_serial query string true "Terminal serial number" example("TRM-001-BUS")
// @Success      200 {object} dto.KeysLoadResponse
// @Failure      400 {object} shareddto.ErrorDefaultResponse "Missing terminal_serial"
// @Failure      403 {object} shareddto.ErrorDefaultResponse "Terminal not found or deactivated"
// @Failure      500 {object} shareddto.ErrorDefaultResponse "Internal server error"
// @Router       /terminal/keys [get]
func (h *TerminalAPIHandler) LoadKeys(c *gin.Context) {
	serial := c.Query("terminal_serial")
	if serial == "" {
		httphelper.JSONError(c, http.StatusBadRequest, errors.New("terminal_serial query parameter is required"))
		return
	}

	resp, err := h.service.LoadKeys(serial)
	if err != nil {
		// Не найден или деактивирован — 403, не 404, чтобы не раскрывать детали
		httphelper.JSONError(c, http.StatusForbidden, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
