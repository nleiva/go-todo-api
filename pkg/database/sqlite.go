package database

import (
	"fmt"
	"log"
	"time"

	"github.com/nleiva/go-todo-api/pkg/app/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLite implements DB for SQLite.
type SQLite struct {
	db *gorm.DB
}

func (m *SQLite) Connect() error {
	dsn := ":memory:"
	return m.connect(dsn)
}

func (m *SQLite) GetDB() *gorm.DB {
	return m.db
}

// ConnectToTestServer connects to the database server without specifying a database
func (m *SQLite) ConnectToTestServer() error {
	dsn := ":memory:"
	return m.connect(dsn)
}

// ConnectToTest connects to the test database
func (m *SQLite) ConnectToTest() error {
	dsn := ":memory:"
	return m.connect(dsn)
}

func (m *SQLite) connect(dsn string) error {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
	})
	if err != nil {
		return fmt.Errorf("[DATABASE]::CONNECTION_ERROR: %w", err)
	}
	m.db = db

	err = m.AutoMigrate()
	if err != nil {
		return fmt.Errorf("[DATABASE]::MIGRATION_ERROR: %w", err)
	}

	return nil
}

func (m *SQLite) AutoMigrate() error {
	return m.db.AutoMigrate(&model.Account{}, &model.Todo{})
}

func (m *SQLite) Disconnect() {
	sqlDB, err := m.db.DB()
	if err != nil {
		log.Println("[DATABASE]::DISCONNECTION_ERROR")
		panic(err)
	}

	sqlDB.Close()
}
