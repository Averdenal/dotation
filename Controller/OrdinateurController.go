package Controller

import (
	"github.com/Averdenal/Dotation/db"
	"github.com/Averdenal/Dotation/logic"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PostOrdinateur(c *gin.Context) {
	nameOrdianteur := c.PostForm("nameOrdianteur")
	codeExpress := c.PostForm("codeExpress")
	os := c.PostForm("os")
	nbSerial := c.PostForm("nbSerial")
	modelOrdianteur := c.PostForm("modelOrdianteur")
	tarif := c.PostForm("tarif")

	o := logic.CreatedOrdinateur(nameOrdianteur, codeExpress, os, nbSerial, modelOrdianteur, tarif)
	db := db.DbConnnect()
	db.Create(&o)
	c.JSON(200, o)
}

func UpdateOrdinateur(c *gin.Context) {
	id := c.Param("id")
	nameOrdianteur := c.PostForm("nameOrdinateur")
	codeExpress := c.PostForm("codeExpress")
	os := c.PostForm("os")
	nbSerial := c.PostForm("nbSerial")
	modelOrdianteur := c.PostForm("modelOrdianteur")
	tarif := c.PostForm("tarif")

	err := logic.UpdateOrdinateur(id, nameOrdianteur, codeExpress, os, nbSerial, modelOrdianteur, tarif)
	if err != nil {
		c.JSON(503, gin.H{
			"err": "error de modification",
		})
	}
	c.JSON(202, "ok")
}

func DeleteOrdinateur(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(503, gin.H{
			"err": err,
		})
		return
	}
	err1 := logic.DeleteOrdianteur(i)
	if err1 != nil {
		c.JSON(503, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(202, "ok")
}

func GetAllOrdinateur(c *gin.Context) {

}

func GetByOrdinateur(c *gin.Context) {

}
