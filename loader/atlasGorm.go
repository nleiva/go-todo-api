package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/nleiva/go-todo-api/pkg/app/model"
)

func main() {
	// Get database type from environment variable, default to sqlite for backward compatibility
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite"
	}

	var stmts string
	var err error

	switch dbType {
	case "mysql":
		stmts, err = gormschema.New("mysql").Load(&model.Account{}, &model.Todo{})
	case "sqlite":
		stmts, err = gormschema.New("sqlite").Load(&model.Account{}, &model.Todo{})
	default:
		fmt.Fprintf(os.Stderr, "unsupported database type: %s. Supported types: mysql, sqlite\n", dbType)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema for %s: %v\n", dbType, err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)
}
