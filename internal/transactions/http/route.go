package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/middleware"
	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/utils"
	"github.com/maxcore25/proga-po-go-transport-service/internal/transactions/service"
)

func RegisterTransactionRoutes(r *gin.RouterGroup, transactionService service.TransactionService, jwtManager *utils.JWTManager) {
	transactionHandler := NewTransactionHandler(transactionService)

	transactionGroup := r.Group("/transactions")
	transactionGroup.GET("", transactionHandler.GetAllTransactions)
	transactionGroup.GET("/:id", transactionHandler.GetTransaction)

	adminProtected := transactionGroup.Group("")
	adminProtected.Use(
		middleware.AuthMiddleware(jwtManager),
		middleware.RoleMiddleware("admin"),
	)
	{
		adminProtected.POST("", transactionHandler.CreateTransaction)
		adminProtected.PATCH("/:id", transactionHandler.UpdateTransactionByID)
		adminProtected.DELETE("/:id", transactionHandler.DeleteTransactionByID)
	}
}
