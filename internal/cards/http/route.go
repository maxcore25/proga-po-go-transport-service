package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/middleware"
	"github.com/maxcore25/proga-po-go-transport-service/internal/cards/service"
	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/utils"
)

func RegisterCardRoutes(r *gin.RouterGroup, cardService service.CardService, jwtManager *utils.JWTManager) {
	cardHandler := NewCardHandler(cardService)

	cardGroup := r.Group("/cards")
	cardGroup.GET("", cardHandler.GetAllCards)
	cardGroup.GET("/:id", cardHandler.GetCard)

	adminProtected := cardGroup.Group("")
	adminProtected.Use(
		middleware.AuthMiddleware(jwtManager),
		middleware.RoleMiddleware("admin"),
	)
	{
		adminProtected.POST("", cardHandler.CreateCard)
		adminProtected.PATCH("/:id", cardHandler.UpdateCardByID)
		adminProtected.DELETE("/:id", cardHandler.DeleteCardByID)
	}
}
