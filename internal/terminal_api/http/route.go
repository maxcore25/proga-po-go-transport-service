package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/proga-po-go-transport-service/internal/terminal_api/service"
)

// RegisterTerminalAPIRoutes регистрирует эндпоинты терминального API.
//
// Эти маршруты не требуют JWT — терминалы идентифицируются
// по serial_number в теле запроса / query-параметре.
//
// Маршруты:
//
//	POST /terminal/authorize  — авторизация платёжной транзакции (5.1)
//	GET  /terminal/keys       — загрузка ключей шифрования MIFARE (5.2)
func RegisterTerminalAPIRoutes(r *gin.RouterGroup, terminalAPIService service.TerminalAPIService) {
	h := NewTerminalAPIHandler(terminalAPIService)

	terminalGroup := r.Group("/terminal")
	{
		terminalGroup.POST("/authorize", h.AuthorizePayment)
		terminalGroup.GET("/keys", h.LoadKeys)
	}
}
