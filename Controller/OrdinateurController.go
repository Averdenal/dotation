package Controller

import (
	"encoding/json"
	"github.com/Averdenal/Dotation/Models"
	"github.com/Averdenal/Dotation/logic"
	"github.com/gin-gonic/gin"
)

func PostOrdinateur(c *gin.Context) {
	var creds Models.CredsOrdinateur
	var o Models.Ordinateur
	err := json.NewDecoder(c.Request.Body).Decode(&creds)
	if err != nil {
		c.AbortWithStatusJSON(400, creds)
	}

	_ = o.Created(creds)
	errSaver := o.Saver()
	if errSaver != nil {
		c.JSON(503, gin.H{
			"err": errSaver,
		})
		return
	}
	c.JSON(200, o)
	return
}

func UpdateOrdinateur(c *gin.Context) {
	var o Models.Ordinateur
	id := c.Param("id")
	nameOrdianteur := c.PostForm("nameOrdinateur")
	codeExpress := c.PostForm("codeExpress")
	nbSerial := c.PostForm("nbSerial")
	idcat := c.PostForm("idcat")
	tarif := c.PostForm("tarif")

	_ = cat.FindById(idcat)
	finderr := o.FindById(id)
	if finderr != nil {
		c.JSON(503, gin.H{
			"err": "find id not found",
		})
		return
	}

	err := logic.UpdateOrdinateur(&o, nameOrdianteur, codeExpress, nbSerial, tarif)
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
	var o Models.Ordinateurs
	err := o.FindAllOrdinateur()
	if err != nil {
		c.JSON(503, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(200, o)
	return
}

func GetOrdinateur(c *gin.Context) {
	id := c.Param("id")
	var o Models.Ordinateur
	err := o.FindById(id)
	if err != nil {
		c.JSON(503, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(200, o)
	return

}
