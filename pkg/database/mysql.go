package database

import (
	"fmt"
	"log"
	"time"

	"github.com/nleiva/go-todo-api/config"
	"github.com/nleiva/go-todo-api/pkg/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL implements DB for MySQL.
type MySQL struct {
	db *gorm.DB
}

func (m *MySQL) Connect() error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_ROOT_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	return m.connect(dsn)
}

func (m *MySQL) GetDB() *gorm.DB {
	return m.db
}

// ConnectToTestServer connects to the database server without specifying a database
func (m *MySQL) ConnectToTestServer() error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.TEST_DB_USER, config.TEST_DB_ROOT_PASSWORD, config.TEST_DB_HOST, config.TEST_DB_PORT, "")
	return m.connect(dsn)
}

// ConnectToTest connects to the test database
func (m *MySQL) ConnectToTest() error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.TEST_DB_USER, config.TEST_DB_ROOT_PASSWORD, config.TEST_DB_HOST, config.TEST_DB_PORT, config.TEST_DB_NAME)
	return m.connect(dsn)
}

func (m *MySQL) connect(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
	})
	if err != nil {
		return fmt.Errorf("[DATABASE]::CONNECTION_ERROR: %w", err)
	}
	m.db = db
	return nil
}

func (m *MySQL) AutoMigrate() error {
	return m.db.AutoMigrate(&model.Account{}, &model.Todo{})
}

func (m *MySQL) Disconnect() {
	sqlDB, err := m.db.DB()
	if err != nil {
		log.Println("[DATABASE]::DISCONNECTION_ERROR")
		panic(err)
	}
	sqlDB.Close()
}
