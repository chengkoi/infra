package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLite SQLite配置
type SQLite struct {
	Path string `yaml:"path"`
}

// NewSQLite 创建SQLite连接
func NewSQLite(config SQLite) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}