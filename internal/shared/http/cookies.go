package httphelper

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	refreshCookieName = "refreshToken"
	refreshCookiePath = "/"
	refreshCookieTTL  = 7 * 24 * time.Hour
)

func refreshTokenCookieSecure() bool {
	return os.Getenv("PORT") == "9000"
}

func refreshTokenCookieSameSite() http.SameSite {
	if !refreshTokenCookieSecure() {
		return http.SameSiteLaxMode
	}
	return http.SameSiteNoneMode
}

// SetRefreshTokenCookie sets a secure, HTTP-only refresh token cookie.
func SetRefreshTokenCookie(c *gin.Context, token string) {
	c.SetSameSite(refreshTokenCookieSameSite())
	c.SetCookie(
		refreshCookieName,
		token,
		int(refreshCookieTTL.Seconds()),
		refreshCookiePath,
		"", // domain (empty = current)
		refreshTokenCookieSecure(),
		true, // httpOnly
	)
}

// GetRefreshTokenFromCookie returns the refresh token stored in the HTTP-only cookie.
func GetRefreshTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie(refreshCookieName)
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", errors.New("empty refresh token cookie")
	}
	return token, nil
}

// ClearRefreshTokenCookie removes the refresh token cookie.
func ClearRefreshTokenCookie(c *gin.Context) {
	c.SetSameSite(refreshTokenCookieSameSite())
	c.SetCookie(
		refreshCookieName,
		"",
		-1, // expire immediately
		refreshCookiePath,
		"",
		refreshTokenCookieSecure(),
		true,
	)
}
