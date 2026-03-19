package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/maxcore25/proga-po-go-transport-service/docs"
)

// @title Transport Card Payment Authorization Service API
// @version 1.0
// @description API documentation for Transport Card Payment Authorization Service.
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.type http
// @securityDefinitions.scheme bearer
// @securityDefinitions.bearerFormat JWT
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer {your_token}" to authorize

func main() {
	config.LoadEnv()

	port := os.Getenv("PORT")

	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "http://127.0.0.1")
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	// Swagger docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(true),
	))

	fmt.Printf("🚀 Dev server started at http://localhost:%s\n", port)
	fmt.Printf("📘 Swagger docs at http://localhost:%s/docs/index.html\n", port)

	r.Run(":" + port)
}
