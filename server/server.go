package server

import (
	"github.com/Averdenal/Dotation/Controller"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

//Server API
func Server() {
	r.Static("/assets", "./assets")

	r.GET("/", Controller.Home)
	
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
