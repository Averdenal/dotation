package Controller

import (
	"github.com/Averdenal/Dotation/logic"
	"github.com/gin-gonic/gin"
)

func NewDotation(c *gin.Context) {
	prenom := c.PostForm("prenom")
	nom := c.PostForm("nom")
	service := c.PostForm("service")

	client, err := logic.NewUser(prenom, nom, service)
	if err != nil {
		c.JSON(503, gin.H{
			"error": err,
		})
	}
	c.JSON(200, client)
}
