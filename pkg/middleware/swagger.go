package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nleiva/go-todo-api/docs"
)

// DynamicSwagger middleware that sets Swagger host based on request
func DynamicSwagger(c *fiber.Ctx) error {
	// Get the host from the request
	host := c.Get("Host")

	// Handle X-Forwarded-Host for reverse proxies
	if forwardedHost := c.Get("X-Forwarded-Host"); forwardedHost != "" {
		host = forwardedHost
	}

	// Update Swagger info with current host
	docs.SwaggerInfo.Host = host

	// Determine scheme based on headers
	scheme := "http"
	if c.Get("X-Forwarded-Proto") == "https" || c.Protocol() == "https" {
		scheme = "https"
	}

	// Update schemes if needed
	if !contains(docs.SwaggerInfo.Schemes, scheme) {
		docs.SwaggerInfo.Schemes = []string{scheme, "http", "https"}
	}

	return c.Next()
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}
