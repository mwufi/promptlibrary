package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"prompt-library/backend/database"
	"prompt-library/backend/handlers"
	"prompt-library/backend/middleware"
	"prompt-library/backend/openai"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/joho/godotenv"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Printf("Warning: Error loading .env file: %v", err)
    }

    // Validate required environment variables
    requiredEnvVars := []string{
        "DATABASE_URL",
        "OPENAI_API_KEY",
    }
    for _, envVar := range requiredEnvVars {
        if os.Getenv(envVar) == "" {
            log.Fatalf("%s environment variable is not set", envVar)
        }
    }

    // Initialize database and run migrations
    log.Println("Initializing database...")
    if err := database.Initialize(os.Getenv("DATABASE_URL")); err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    log.Println("Database initialization completed")

    // Initialize OpenAI service
    openaiService := openai.NewService(os.Getenv("OPENAI_API_KEY"))

    // Create Fiber app with custom config
    app := fiber.New(fiber.Config{
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
        BodyLimit:    1 * 1024 * 1024, // 1MB
    })

    // Add security middleware
    middleware.SecurityMiddleware(app)

    // Add recover middleware to handle panics
    app.Use(recover.New())

    // Add logger
    app.Use(logger.New(logger.Config{
        Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
    }))

    // Add CORS middleware with strict configuration
    corsOrigins := os.Getenv("CORS_ORIGINS")
    if corsOrigins == "" {
        corsOrigins = "http://localhost:3000" // Default to local development
    }
    app.Use(cors.New(cors.Config{
        AllowOrigins:     corsOrigins,
        AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
        AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
        ExposeHeaders:    "Content-Length",
        AllowCredentials: true,
        MaxAge:           300,
    }))

    // Add middleware to inject OpenAI service
    app.Use(func(c *fiber.Ctx) error {
        c.Locals("openai", openaiService)
        return c.Next()
    })

    // Setup routes with timeouts
    api := app.Group("/api")
    
    // Prompts routes with timeouts
    api.Get("/prompts", timeout.New(handlers.HandleGetPrompts, 30*time.Second))
    api.Post("/prompts", timeout.New(handlers.HandleCreatePrompt, 30*time.Second))
    api.Get("/prompts/:id", timeout.New(handlers.HandleGetPrompt, 30*time.Second))
    api.Post("/prompts/:id/vote", timeout.New(handlers.HandleVotePrompt, 30*time.Second))
    api.Post("/prompts/:id/run", timeout.New(handlers.HandleRunPrompt, 60*time.Second))

    // Start server in a goroutine
    go func() {
        port := os.Getenv("PORT")
        if port == "" {
            port = "8000"
        }
        log.Printf("Starting server on port %s...", port)
        if err := app.Listen(":" + port); err != nil {
            log.Printf("Server error: %v", err)
        }
    }()

    // Setup graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    log.Println("Shutting down server...")
    
    // Give outstanding requests 5 seconds to complete
    if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
        log.Printf("Error during server shutdown: %v", err)
    }
}
