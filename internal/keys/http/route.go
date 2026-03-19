package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/middleware"
	"github.com/maxcore25/proga-po-go-transport-service/internal/keys/service"
	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/utils"
)

func RegisterKeyRoutes(r *gin.RouterGroup, keyService service.KeyService, jwtManager *utils.JWTManager) {
	keyHandler := NewKeyHandler(keyService)

	keyGroup := r.Group("/keys")
	keyGroup.GET("", keyHandler.GetAllKeys)
	keyGroup.GET("/:id", keyHandler.GetKey)

	adminProtected := keyGroup.Group("")
	adminProtected.Use(
		middleware.AuthMiddleware(jwtManager),
		middleware.RoleMiddleware("admin"),
	)
	{
		adminProtected.POST("", keyHandler.CreateKey)
		adminProtected.PATCH("/:id", keyHandler.UpdateKeyByID)
		adminProtected.DELETE("/:id", keyHandler.DeleteKeyByID)
	}
}
