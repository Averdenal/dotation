package Controller

import (
	"fmt"
	"github.com/Averdenal/Dotation/Models"
	"github.com/Averdenal/Dotation/logic"
	"github.com/gin-gonic/gin"
)

func PostOrdinateur(c *gin.Context) {
	nameOrdinateur := c.PostForm("nameOrdinateur")
	codeExpress := c.PostForm("codeExpress")
	os := c.PostForm("os")
	nbSerial := c.PostForm("nbSerial")
	modelOrdianteur := c.PostForm("modelOrdianteur")
	tarif := c.PostForm("tarif")

	o, _ := logic.CreatedOrdinateur(nameOrdinateur, codeExpress, os, nbSerial, modelOrdianteur, tarif)
	fmt.Println(o)
	err := o.Saver()
	if err != nil {
		c.JSON(503, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(200, o)
	return
}

func UpdateOrdinateur(c *gin.Context) {
	id := c.Param("id")
	nameOrdianteur := c.PostForm("nameOrdinateur")
	codeExpress := c.PostForm("codeExpress")
	os := c.PostForm("os")
	nbSerial := c.PostForm("nbSerial")
	modelOrdianteur := c.PostForm("modelOrdianteur")
	tarif := c.PostForm("tarif")

	var o Models.Ordinateur
	finderr := o.FindById(id)
	if finderr != nil {
		c.JSON(503, gin.H{
			"err": "find id not found",
		})
		return
	}

	fmt.Println(o)
	err := logic.UpdateOrdinateur(&o, nameOrdianteur, codeExpress, os, nbSerial, modelOrdianteur, tarif)
	if err != nil {
		c.JSON(503, gin.H{
			"err": "error logic de modification",
		})
		return
	}
	saverErr := o.Saver()
	if finderr != nil {
		c.JSON(503, gin.H{
			"err": saverErr,
		})
	}
	c.JSON(202, "ok")
}

func DeleteOrdinateur(c *gin.Context) {
	id := c.Param("id")
	var o Models.Ordinateur
	_ = o.FindById(id)
	err := o.Delete()
	if err != nil {
		c.JSON(503, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(202, "ok")
}

func GetAllOrdinateur(c *gin.Context) {

}

func GetOrdinateur(c *gin.Context) {

}
