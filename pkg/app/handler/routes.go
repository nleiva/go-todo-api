package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nleiva/go-todo-api/config"
	"github.com/nleiva/go-todo-api/pkg/middleware"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// New registers all routes for the application
func (h *Handler) RegisterRoutes(app *fiber.App) {
	h.RegisterHyperMediaRoutes(app)
	h.RegisterApiRoutes(app)
}

// RegisterHyperMediaRoutes registers all routes for the application that are used for rendering views
//
// Good Read for Hypermedia-Driven Applications: https://hypermedia.systems/json-data-apis/
func (h *Handler) RegisterHyperMediaRoutes(app *fiber.App) {
	app.Static("/js", config.ROOT_PATH+"/pkg/view/js")

	app.Get("/", middleware.LoadAuth, h.VIndex)

	app.Get("/login", h.VLogin)
	app.Post("/login", h.VLoginPost)
	app.Post("/logout", h.VLogout)

	app.Get("/register", h.VRegister)
	app.Post("/register", h.VRegisterPost)

	app.Get("/todos", middleware.Pagination, middleware.Protected, middleware.Pagination, h.VTodosIndex)
	app.Post("/todos", middleware.Protected, h.VTodosCreate)
	app.Put("/todos/:id/complete", middleware.Protected, h.VTodosComplete)
	app.Delete("/todos/:id", middleware.Protected, h.VTodosDelete)
}

func (h *Handler) RegisterApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from base api path")
	})

	// Swagger UI documentation
	api.Get("/docs/*", fiberSwagger.WrapHandler)

	// Redoc documentation
	api.Get("/redoc", func(c *fiber.Ctx) error {
		return c.Type("html").SendString(`
<!DOCTYPE html>
<html>
<head>
    <title>API Documentation</title>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
    <style>
      body { margin: 0; padding: 0; }
    </style>
</head>
<body>
    <redoc spec-url='/api/docs/doc.json'></redoc>
    <script src="https://cdn.jsdelivr.net/npm/redoc@latest/bundles/redoc.standalone.js"></script>
</body>
</html>`)
	})

	auth := api.Group("/auth")
	auth.Put("/login", h.Login)
	auth.Post("/register", h.Register)
	auth.Put("/refresh", middleware.Protected, h.Refresh)
	auth.Get("/me", middleware.Protected, h.Me)

	auth.Put("/jwk-rotate", middleware.AllowedIps, h.RotateJWK)

	accounts := api.Group("/accounts")
	accounts.Get("/", middleware.Protected, middleware.Pagination, h.GetAccounts)
	accounts.Get("/:id", middleware.Protected, h.GetAccount)

	todos := api.Group("/todos")
	todos.Get("/", middleware.Protected, middleware.Pagination, h.GetTodos)
	todos.Get("/csv", middleware.Protected, h.ExportCSVTodos)
	todos.Get("/:id", middleware.Protected, h.GetTodo)
	todos.Post("/", middleware.Protected, h.CreateTodo)
	todos.Post("/csv", middleware.Protected, h.ImportCSVTodos)
	todos.Put("/:id", middleware.Protected, h.UpdateTodo)
	todos.Delete("/:id", middleware.Protected, h.DeleteTodo)

	todos.Post("/random", middleware.Protected, h.CreateRandomTodo)
}
