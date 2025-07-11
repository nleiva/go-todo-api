package model_test

import (
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nleiva/go-todo-api/pkg/app"
	"github.com/nleiva/go-todo-api/test"
	"gorm.io/gorm"
)

var testApp *fiber.App
var DB *gorm.DB

func TestMain(m *testing.M) {
	// Setup
	DB = test.Setup()
	testApp = app.New(DB)

	// Run tests
	code := m.Run()

	// Teardown
	app.Shutdown(testApp)
	test.Teardown(DB)

	os.Exit(code)
}
