package main

import (
	"fmt"

	"github.com/Averdenal/Dotation/logic"
	"github.com/Averdenal/Dotation/server"
)

func main() {
	server.Server()
	c, err := logic.NewUser("amaury", "verdenal", "Informatique")
	if err != nil {
		fmt.Printf("error = %v", err)
	}
	logic.AddOrdinateur(&c, "BGSA-O-145-prt", "48575214711", "Windows 7", "RF584W2", "7480", 1100.50)
	app, _ := logic.CreatedApplication("PMAD", "Dameware")
	logic.AddApplication(&c, app)

	app, _ = logic.CreatedApplication("Messagerie", "Outlook")
	logic.AddApplication(&c, app)

	mat, _ := logic.CreatedMateriel( 200.5, "er54875421-58","Ecran","22p")
	logic.AddMateriel(&c, mat)
	fmt.Printf("add ordinateur \n %v", c)
}
