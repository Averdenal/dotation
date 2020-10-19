package db

import (
	"github.com/Averdenal/Dotation/Models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConnnect() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return db
}

func Actualisation() {
	db := DbConnnect()
	db.AutoMigrate(&Models.User{})
	db.AutoMigrate(&Models.Application{})
	db.AutoMigrate(&Models.Ordinateur{})
	db.AutoMigrate(&Models.Materiel{})
	db.AutoMigrate(&Models.Cat{})
	db.AutoMigrate(&Models.Dotation{})
}
