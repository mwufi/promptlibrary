package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"time"
)

// SecurityMiddleware adds security-related middleware to the application
func SecurityMiddleware(app *fiber.App) {
	// Add security headers
	app.Use(helmet.New())

	// Recover from panics
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Rate limiting
	app.Use(limiter.New(limiter.Config{
		Max:        100,                // 100 requests
		Expiration: 1 * time.Minute,    // per minute
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limit by IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	}))
}
