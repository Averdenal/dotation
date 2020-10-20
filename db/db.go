package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConnnect() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	return db
}
