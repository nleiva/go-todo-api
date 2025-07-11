package handler_test

import (
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nleiva/go-todo-api/pkg/app"
	"github.com/nleiva/go-todo-api/test"
	"gorm.io/gorm"
)

var App *fiber.App
var DB *gorm.DB

func TestMain(m *testing.M) {
	// Setup
	DB = test.Setup()
	if DB == nil {
		panic("Failed to setup database")
	}
	App = app.New(DB)
	if App == nil {
		panic("Failed to create app")
	}

	// Run tests
	code := m.Run()

	// Teardown
	app.Shutdown(App)
	test.Teardown(DB)

	os.Exit(code)
}
