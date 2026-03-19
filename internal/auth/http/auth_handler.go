package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/service"
	httphelper "github.com/maxcore25/proga-po-go-transport-service/internal/shared/http"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

// Register godoc
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body dto.RegisterRequest true "User registration data"
// @Success 200 {object} dto.AuthTokens
// @Failure 400 {object} gin.H
// @Failure 409 {object} gin.H
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if !httphelper.BindJSON(c, &req) {
		return
	}

	tokens, err := h.service.Register(req)
	if err != nil {
		if err.Error() == "user with this email already exists" {
			httphelper.JSONError(c, http.StatusConflict, err)
		} else {
			httphelper.JSONError(c, http.StatusBadRequest, err)
		}
		return
	}

	httphelper.SetRefreshTokenCookie(c, tokens.RefreshToken)
	httphelper.JSONAccessToken(c, tokens.AccessToken)
}

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "User credentials"
// @Success 200 {object} dto.AuthTokens
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if !httphelper.BindJSON(c, &req) {
		return
	}

	tokens, err := h.service.Login(req)
	if err != nil {
		httphelper.JSONError(c, http.StatusUnauthorized, err)
		return
	}

	httphelper.SetRefreshTokenCookie(c, tokens.RefreshToken)
	httphelper.JSONAccessToken(c, tokens.AccessToken)
}

// Refresh godoc
// @Summary Refresh access token
// @Tags Auth
// @Produce json
// @Success 200 {object} dto.AuthTokens
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	// Extract token from httpOnly cookie
	refreshToken, err := httphelper.GetRefreshTokenFromCookie(c)
	if err != nil {
		httphelper.JSONError(c, http.StatusBadRequest, errors.New("refresh token cookie missing"))
		return
	}

	// Attempt to refresh tokens
	tokens, err := h.service.Refresh(refreshToken)
	if err != nil {
		httphelper.JSONError(c, http.StatusUnauthorized, err)
		return
	}

	// Rotate refresh token cookie
	httphelper.SetRefreshTokenCookie(c, tokens.RefreshToken)

	// Return access token in JSON body
	httphelper.JSONAccessToken(c, tokens.AccessToken)
}

// Logout godoc
// @Summary Logout user (invalidate refresh token)
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	// Read refresh token from HTTP-only cookie
	refreshToken, err := httphelper.GetRefreshTokenFromCookie(c)
	if err != nil {
		// cookie missing or invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token cookie not found"})
		return
	}

	// Try to remove refresh token from store (service should handle missing token gracefully)
	if err := h.service.Logout(refreshToken); err != nil {
		// return 500 only on internal error; if token was already gone, service can return nil
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Clear cookie on client
	httphelper.ClearRefreshTokenCookie(c)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
