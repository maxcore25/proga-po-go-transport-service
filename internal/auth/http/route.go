package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/middleware"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/service"
	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/utils"
)

func RegisterAuthRoutes(r *gin.RouterGroup, userService service.UserService, authService service.AuthService, jwtManager *utils.JWTManager) {
	userHandler := NewUserHandler(userService)
	authHandler := NewAuthHandler(authService)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/refresh", authHandler.Refresh)
		authGroup.POST("/logout", authHandler.Logout)
	}

	userGroup := r.Group("/users")
	userGroup.GET("", userHandler.GetAllUsers)
	userGroup.GET("/:id", userHandler.GetUser)

	protected := userGroup.Group("")
	protected.Use(middleware.AuthMiddleware(jwtManager))
	{
		protected.POST("", userHandler.CreateUser)
		protected.GET("/me", userHandler.GetCurrentUser)
		protected.PATCH("/:id", userHandler.UpdateUserByID)
		protected.DELETE("/:id", userHandler.DeleteUserByID)
	}
}
