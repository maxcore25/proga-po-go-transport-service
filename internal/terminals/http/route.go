package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/middleware"
	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/utils"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminals/service"
)

func RegisterTerminalRoutes(r *gin.RouterGroup, terminalService service.TerminalService, jwtManager *utils.JWTManager) {
	terminalHandler := NewTerminalHandler(terminalService)

	terminalGroup := r.Group("/terminals")
	terminalGroup.GET("", terminalHandler.GetAllTerminals)
	terminalGroup.GET("/:id", terminalHandler.GetTerminal)

	adminProtected := terminalGroup.Group("")
	adminProtected.Use(
		middleware.AuthMiddleware(jwtManager),
		middleware.RoleMiddleware("admin"),
	)
	{
		adminProtected.POST("", terminalHandler.CreateTerminal)
		adminProtected.PATCH("/:id", terminalHandler.UpdateTerminalByID)
		adminProtected.DELETE("/:id", terminalHandler.DeleteTerminalByID)
	}
}
