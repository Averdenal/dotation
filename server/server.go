package server

import (
	"log"

	"github.com/Averdenal/Dotation/Controller"
	"github.com/gin-gonic/gin"

	"html/template"
)

var r = gin.Default()

//Server API
func Server() {
	r.Static("/assets", "./assets")
	r.GET("/", demo)
	demandes := r.Group("/Demandes")
	{
		//vos demandes
		demandes.GET("/", Controller.DemandeByUser)
		//Créer une demande
		demandes.POST("/NewUser", Controller.NewUser)

		demandes.POST("/", Controller.NewDemande)

		demandes.GET("/user/:idUser", Controller.DemandeForUser)

		demandes.GET("/Service/:idService", Controller.NewDotation)

	}

	r.Run()
}
func demo(c *gin.Context) {
	p := Page{"Titre de ma page", []string{"Item 1", "Item 2", "Item 3"}}

	t := template.New("home page")

	t = template.Must(t.ParseFiles("./Template/layout.tmpl", "./Template/content.tmpl"))

	err := t.ExecuteTemplate(c.Writer, "layout", p)
	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

type Page struct {
	Title    string
	Articles []string
}
