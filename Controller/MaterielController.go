package Controller

import (
	"github.com/Averdenal/Dotation/Models"
	"github.com/Averdenal/Dotation/logic"
	"github.com/gin-gonic/gin"
)

//GetAllMateriel trouve toutes les entrées de matériel sauf les DELETE
func GetAllMateriel(c *gin.Context) {
	var m Models.Materiels
	err := m.FindAllMateriel()
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.JSON(200, m)
	return
}

func GetMateriel(c *gin.Context) {
	id := c.Param("id")
	var m Models.Materiel
	err := m.FindById(id)
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.JSON(200, m)
	return

}

func PostMateriel(c *gin.Context) {
	t := c.PostForm("tarif")
	nbs := c.PostForm("nbserial")
	n := c.PostForm("name")
	m := c.PostForm("model")

	cat := Models.Cat{
		Type:   n,
		Models: m,
	}

	mat, err := logic.CreatedMateriel(t, nbs, cat)
	if err != nil {
		c.JSON(503, err)
		return
	}
	err1 := mat.Saver()
	if err1 != nil {
		c.JSON(503, err1)
		return
	}
	c.JSON(202, mat)
	return

}

func DeleteMateriel(c *gin.Context) {
	i := c.Param("id")
	var m Models.Materiel
	err := m.FindById(i)
	if err != nil {
		c.JSON(503, err)
		return
	}
	errDelete := m.Delete()
	if errDelete != nil {
		c.JSON(503, errDelete)
		return
	}
	c.JSON(202, "suppr ok")
}

func UpdateMateriel(c *gin.Context) {
	id := c.Param("id")
	t := c.PostForm("tarif")
	nbs := c.PostForm("nbserial")
	n := c.PostForm("name")
	m := c.PostForm("model")
	var mat Models.Materiel
	err := mat.FindById(id)
	if err != nil {
		c.JSON(503, err)
		return
	}
	cat := Models.Cat{
		Type:   n,
		Models: m,
	}
	logic.UpdateMateriel(&mat, t, nbs, cat)
	errSaver := mat.Saver()
	if errSaver != nil {
		c.JSON(503, errSaver)
		return
	}
	c.JSON(202, mat)
}
