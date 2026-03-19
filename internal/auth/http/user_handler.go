package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/mapper"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/service"
	httphelper "github.com/maxcore25/proga-po-go-transport-service/internal/shared/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// CreateUser godoc
// @Summary Create user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body dto.RegisterRequest true "New user"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.RegisterRequest

	if !httphelper.BindJSON(c, &req) {
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.NewUserResponse(user)

	c.JSON(http.StatusCreated, resp)
}

// GetCurrentUser godoc
// @Summary Get current user
// @Description Returns the authenticated user
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /users/me [get]
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		return
	}
	id, ok := userID.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id type"})
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return

	}

	resp := mapper.NewUserResponse(user)

	c.JSON(http.StatusOK, resp)
}

// GetUser godoc
// @Summary Get user by ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID (uuid)"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	resp := mapper.NewUserResponse(user)

	c.JSON(http.StatusOK, resp)
}

// GetAllUsers godoc
// @Summary Get all users
// @Tags Users
// @Produce json
// @Param role query string false "User role filter. One of: client, admin, tutor"
// @Success 200 {array} dto.UserResponse
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	// * v1 (without filters)
	// users, err := h.service.GetAllUsers()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// resp := make([]*dto.UserResponse, len(users))
	// for i, user := range users {
	// 	resp[i] = mapper.NewUserResponse(user)
	// }
	// c.JSON(http.StatusOK, resp)

	// * v2 (with filters)
	var filter dto.UserFilter

	// Bind query params like ?role=tutor
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filters"})
		return
	}

	users, err := h.service.GetUsers(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		resp[i] = mapper.NewUserResponse(user)
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUserByID godoc
// @Summary Update user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID (uuid)"
// @Param user body dto.UpdateUserRequest true "User update data"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [patch]
func (h *UserHandler) UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	var updateData dto.UpdateUserRequest
	if !httphelper.BindJSON(c, &updateData) {
		return
	}
	if err := h.service.UpdateUserByID(id, updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

// DeleteUserByID godoc
// @Summary Delete user by ID
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID (uuid)"
// @Success 204 {object} nil
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	if err := h.service.DeleteUserByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
