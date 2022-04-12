package database

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
)

func NewDatabase(connectionSql string) *gorm.DB {
	db, err := gorm.Open(sqlserver.Open(connectionSql), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}
	return db
}
