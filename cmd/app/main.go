package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/pressly/goose"
	"gorm.io/gorm"

	authHttp "github.com/maxcore25/proga-po-go-transport-service/internal/auth/http"
	authRepo "github.com/maxcore25/proga-po-go-transport-service/internal/auth/repository"
	authService "github.com/maxcore25/proga-po-go-transport-service/internal/auth/service"

	cardHttp "github.com/maxcore25/proga-po-go-transport-service/internal/cards/http"
	cardRepo "github.com/maxcore25/proga-po-go-transport-service/internal/cards/repository"
	cardService "github.com/maxcore25/proga-po-go-transport-service/internal/cards/service"

	keyHttp "github.com/maxcore25/proga-po-go-transport-service/internal/keys/http"
	keyRepo "github.com/maxcore25/proga-po-go-transport-service/internal/keys/repository"
	keyService "github.com/maxcore25/proga-po-go-transport-service/internal/keys/service"

	terminalHttp "github.com/maxcore25/proga-po-go-transport-service/internal/terminals/http"
	terminalRepo "github.com/maxcore25/proga-po-go-transport-service/internal/terminals/repository"
	terminalService "github.com/maxcore25/proga-po-go-transport-service/internal/terminals/service"

	transactionHttp "github.com/maxcore25/proga-po-go-transport-service/internal/transactions/http"
	transactionRepo "github.com/maxcore25/proga-po-go-transport-service/internal/transactions/repository"
	transactionService "github.com/maxcore25/proga-po-go-transport-service/internal/transactions/service"

	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/config"
	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/utils"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/maxcore25/proga-po-go-transport-service/docs"
)

// @title Transport Card Payment Authorization Service API
// @version 1.0
// @description API documentation for Transport Card Payment Authorization Service.
// @host localhost:8888
// @BasePath /api/v1
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

	jwtAccessSecret := os.Getenv("JWT_ACCESS_SECRET")
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	accessExpStr := os.Getenv("JWT_ACCESS_EXPIRATION")
	refreshExpStr := os.Getenv("JWT_REFRESH_EXPIRATION")

	accessExp, err := utils.ParseDuration(accessExpStr)
	if err != nil {
		log.Fatalf("Invalid JWT_ACCESS_EXPIRATION: %v", err)
	}
	refreshExp, err := utils.ParseDuration(refreshExpStr)
	if err != nil {
		log.Fatalf("Invalid JWT_REFRESH_EXPIRATION: %v", err)
	}

	// Создаём директорию для БД если её нет
	dbPath := "app.db"
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("❌ Failed to create db dir: %v", err)
	}

	// --- 1. Открываем sql.DB (для goose) ---
	dbSQL, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("❌ Failed to open database: %v", err)
	}

	// SQLite работает лучше с одним writer-соединением
	dbSQL.SetMaxOpenConns(1)
	dbSQL.SetMaxIdleConns(1)

	if err := dbSQL.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping database: %v", err)
	}

	// --- 2. Применяем миграции ---
	goose.SetDialect("sqlite3")
	if err := goose.Up(dbSQL, "./internal/migrations"); err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	// --- 3. Поднимаем GORM поверх той же БД ---
	db, err := gorm.Open(sqlite.Dialector{
		DSN: dbPath,
	}, &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect database: %v", err)
	}

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

	jwtManager := utils.NewJWTManager(jwtAccessSecret, jwtRefreshSecret)
	jwtManager.AccessTokenTTL = accessExp
	jwtManager.RefreshTokenTTL = refreshExp

	userRepo := authRepo.NewUserRepository(db)
	userService := authService.NewUserService(userRepo)
	refreshTokenRepo := authRepo.NewRefreshTokenRepository(db)
	authService := authService.NewAuthService(userRepo, refreshTokenRepo, jwtManager)
	cardsRepo := cardRepo.NewCardRepository(db)
	cardsService := cardService.NewCardService(cardsRepo)
	keysRepo := keyRepo.NewKeyRepository(db)
	keysService := keyService.NewKeyService(keysRepo)
	terminalsRepo := terminalRepo.NewTerminalRepository(db)
	terminalsService := terminalService.NewTerminalService(terminalsRepo)
	transactionsRepo := transactionRepo.NewTransactionRepository(db)
	transactionsService := transactionService.NewTransactionService(transactionsRepo)

	// Register routes
	api := r.Group("/api/v1")
	{
		authHttp.RegisterAuthRoutes(api, userService, authService, jwtManager)
		cardHttp.RegisterCardRoutes(api, cardsService, jwtManager)
		keyHttp.RegisterKeyRoutes(api, keysService, jwtManager)
		terminalHttp.RegisterTerminalRoutes(api, terminalsService, jwtManager)
		transactionHttp.RegisterTransactionRoutes(api, transactionsService, jwtManager)
	}

	// Swagger docs
	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(true),
	))

	fmt.Printf("🚀 Dev server started at http://localhost:%s\n", port)
	fmt.Printf("📘 Swagger docs at http://localhost:%s/api/v1/swagger/index.html\n", port)

	r.Run(":" + port)
}
