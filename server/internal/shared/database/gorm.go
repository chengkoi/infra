package database

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// Config 数据库配置
type Config struct {
	Driver string `yaml:"driver"`
	MySQL  MySQL  `yaml:"mysql"`
	PostgreSQL
	SQLite SQLite `yaml:"sqlite"`
}

// Init 初始化数据库
func Init(config Config) error {
	var err error
	var db *gorm.DB

	switch config.Driver {
	case "mysql":
		db, err = NewMySQL(config.MySQL)
	case "postgres":
		db, err = NewPostgreSQL(config.PostgreSQL)
	case "sqlite":
		db, err = NewSQLite(config.SQLite)
	default:
		return fmt.Errorf("unsupported database driver: %s", config.Driver)
	}

	if err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}

	DB = db
	return nil
}