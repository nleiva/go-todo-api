package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// NoCacheForDevelopment adds cache control headers to prevent caching during development
func NoCacheForDevelopment(c *fiber.Ctx) error {
	// Only apply no-cache headers for HTML pages and development
	path := c.Path()

	// For HTML pages and API endpoints, prevent caching
	if strings.HasSuffix(path, "/") ||
		strings.Contains(path, "/api/") ||
		!strings.Contains(path, ".") {
		c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
	}

	// For static assets (JS, CSS), use short cache
	if strings.HasSuffix(path, ".js") ||
		strings.HasSuffix(path, ".css") {
		c.Set("Cache-Control", "max-age=60") // 1 minute cache
	}

	return c.Next()
}

// CacheBusting adds cache busting headers for development
func CacheBusting(c *fiber.Ctx) error {
	// Add ETag based on current time for development
	c.Set("ETag", `"`+c.Get("X-Request-ID")+`"`)

	return c.Next()
}
