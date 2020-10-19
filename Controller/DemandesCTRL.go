package Controller

import (
	"fmt"
	"github.com/Averdenal/Dotation/db"
	"github.com/Averdenal/Dotation/logic"
	"github.com/gin-gonic/gin"
)

//Demandes GET All demandes
func Demandes(c *gin.Context) {

}

//UserDemande GET Demande par un User (:id)
func DemandeByUser(c *gin.Context) {

}

//UserDemande GET Demande pour un User (:id)
func DemandeForUser(c *gin.Context) {

}

//ServiceDemandes GET demandes par service (:id)
func ServiceDemandes(c *gin.Context) {

}

//NewDemande POST Création d'une demande
func NewDemande(c *gin.Context) {

}

func NewUser(c *gin.Context) {
	prenom := c.PostForm("prenom")
	nom := c.PostForm("nom")
	service := c.PostForm("service")

	client, err := logic.NewUser(prenom, nom, service)
	if err != nil {
		fmt.Printf("error = %v", err)
	}

	db := db.DbConnnect()
	db.Create(&client)
	if err != nil {
		c.JSON(503, gin.H{
			"error": err,
		})
	}
	c.JSON(200, client)

}

func NewApplication(c *gin.Context) {
	name := c.PostForm("name")
	version := c.PostForm("version")

	app, _ := logic.CreatedApplication(name, version)
	db := db.DbConnnect()
	db.Create(&app)
	c.JSON(200, app)
}
