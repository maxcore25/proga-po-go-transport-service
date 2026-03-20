package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/mapper"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/service"
	shareddto "github.com/maxcore25/proga-po-go-transport-service/internal/shared/dto"
	httphelper "github.com/maxcore25/proga-po-go-transport-service/internal/shared/http"
)

type KeyHandler struct {
	service service.KeyService
}

func NewKeyHandler(s service.KeyService) *KeyHandler {
	return &KeyHandler{service: s}
}

// CreateKey godoc
// @Summary Create key
// @Tags Keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param key body dto.CreateKeyRequest true "New key"
// @Success 201 {object} dto.KeyResponse
// @Failure 400 {object} shareddto.ErrorDefaultResponse
// @Failure 500 {object} shareddto.ErrorDefaultResponse
// @Router /keys [post]
func (h *KeyHandler) CreateKey(c *gin.Context) {
	var req dto.CreateKeyRequest

	if !httphelper.BindJSON(c, &req) {
		return
	}

	key, err := h.service.CreateKey(req)
	if err != nil {
		httphelper.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	resp := mapper.NewKeyResponse(key)

	c.JSON(http.StatusCreated, resp)
}

// GetKey godoc
// @Summary Get key by ID
// @Tags Keys
// @Produce json
// @Param id path string true "Key ID (uuid)"
// @Success 200 {object} dto.KeyResponse
// @Failure 400 {object} shareddto.ErrorDefaultResponse
// @Failure 404 {object} shareddto.ErrorDefaultResponse
// @Router /keys/{id} [get]
func (h *KeyHandler) GetKey(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		httphelper.JSONError(c, http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}

	key, err := h.service.GetKey(id)
	if err != nil {
		httphelper.JSONError(c, http.StatusNotFound, errors.New("key not found"))
		return
	}

	resp := mapper.NewKeyResponse(key)

	c.JSON(http.StatusOK, resp)
}

// GetAllKeys godoc
// @Summary Get all keys
// @Tags Keys
// @Produce json
// @Success 200 {array} dto.KeyResponse
// @Router /keys [get]
func (h *KeyHandler) GetAllKeys(c *gin.Context) {
	keys, err := h.service.GetAllKeys()
	if err != nil {
		httphelper.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	resp := make([]*dto.KeyResponse, len(keys))
	for i, key := range keys {
		resp[i] = mapper.NewKeyResponse(key)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateKeyByID godoc
// @Summary Update key by ID
// @Tags Keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Key ID (uuid)"
// @Param key body dto.UpdateKeyRequest true "Key update data"
// @Success 200 {object} shareddto.MessageDefaultResponse
// @Failure 400 {object} shareddto.ErrorDefaultResponse
// @Failure 404 {object} shareddto.ErrorDefaultResponse
// @Failure 500 {object} shareddto.ErrorDefaultResponse
// @Router /keys/{id} [patch]
func (h *KeyHandler) UpdateKeyByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		httphelper.JSONError(c, http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}
	var updateData dto.UpdateKeyRequest
	if !httphelper.BindJSON(c, &updateData) {
		return
	}
	if err := h.service.UpdateKeyByID(id, updateData); err != nil {
		httphelper.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, shareddto.MessageDefaultResponse{Message: "key updated successfully"})
}

// DeleteKeyByID godoc
// @Summary Delete key by ID
// @Tags Keys
// @Produce json
// @Security BearerAuth
// @Param id path string true "Key ID (uuid)"
// @Success 204 {object} nil
// @Failure 400 {object} shareddto.ErrorDefaultResponse
// @Failure 404 {object} shareddto.ErrorDefaultResponse
// @Failure 500 {object} shareddto.ErrorDefaultResponse
// @Router /keys/{id} [delete]
func (h *KeyHandler) DeleteKeyByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		httphelper.JSONError(c, http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}
	if err := h.service.DeleteKeyByID(id); err != nil {
		httphelper.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
