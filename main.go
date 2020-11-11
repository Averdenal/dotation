package main

import (
	"github.com/Averdenal/Dotation/Models"
	"github.com/Averdenal/Dotation/db"
	"github.com/Averdenal/Dotation/server"
)

func main() {
	data := db.DbConnnect()
	data.AutoMigrate(&Models.User{})
	data.AutoMigrate(&Models.Application{})
	data.AutoMigrate(&Models.Cat{})
	data.AutoMigrate(&Models.Materiel{})
	data.AutoMigrate(&Models.Ordinateur{})
	data.AutoMigrate(&Models.Service{})
	server.Server()
}
