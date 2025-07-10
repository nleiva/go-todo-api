package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/nleiva/go-todo-api/api"
	"github.com/nleiva/go-todo-api/pkg/app"
	"github.com/nleiva/go-todo-api/pkg/database"

	"gorm.io/gorm"
)

// @title           fiber-api
// @version         1.0
// @BasePath  /api

type backend interface {
	Connect() error
	GetDB() *gorm.DB
}

func run() error {
	PORT := "3000"
	DB := "sqlite"

	var b backend
	switch DB {
	case "mysql":
		b = &database.MySQL{}
	case "sqlite":
		b = &database.SQLite{}
	}

	err := b.Connect()
	if err != nil {
		return fmt.Errorf("cannot connect to database: %w", err)
	}

	app := app.New(b.GetDB())

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start the Fiber app in a goroutine
	go func() {
		if err := app.Listen(":" + PORT); err != nil {
			log.Fatalf("Fiber server error: %v", err)
		}
	}()

	// Block until a signal is received
	<-quit
	log.Println("shutting down Fiber application...")

	// Perform graceful shutdown
	if err := app.ShutdownWithTimeout(1 * time.Second); err != nil {
		log.Fatalf("Fiber shutdown error: %v", err)
	}

	log.Println("Fiber application gracefully shut down.")

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error running: %s\n", err)
		os.Exit(1)
	}
}
