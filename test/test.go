package test

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/nleiva/go-todo-api/config"
	"github.com/nleiva/go-todo-api/pkg/database"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	if !config.IS_TEST {
		panic("[New]::IS_TEST is not true")
	}
	if config.ROOT_PATH == "" {
		panic("[New]::ROOT_PATH is empty. Please set the environment variable GTA_ROOT_PATH to the root path of the project (See makefile:test)")
	}

	os.MkdirAll(config.TEST_FILE_PATH, fs.ModePerm)

	// Use SQLite for testing (in-memory)
	db := &database.SQLite{}
	err := db.Connect()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to test database: %v", err))
	}

	dbInstance := db.GetDB()
	if dbInstance == nil {
		panic("Database instance is nil after connection")
	}

	return dbInstance
}

func Teardown(db *gorm.DB) {
	os.RemoveAll(config.TEST_FILE_PATH)

	if db != nil {
		sqlDB, _ := db.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}
}

func ClearTables(db *gorm.DB, tables []string) {
	clearTables(tables, db)
}

func ClearAllTables(db *gorm.DB) {
	// For SQLite, we can just delete all records from tables
	db.Exec("DELETE FROM todos")
	db.Exec("DELETE FROM accounts")
}

func clearTables(tables []string, db *gorm.DB) {
	for _, table := range tables {
		db.Exec("DELETE FROM " + table)
	}
}
