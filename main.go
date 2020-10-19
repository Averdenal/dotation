package main

import (
	"github.com/Averdenal/Dotation/db"
	"github.com/Averdenal/Dotation/server"
)

func main() {
	db.Actualisation()
	server.Server()

}
