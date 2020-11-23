package Controller

import (
	"github.com/Averdenal/Dotation/Models"
	"github.com/Averdenal/Dotation/logic"
	"github.com/gin-gonic/gin"
)

var cat Models.Cat
var cats Models.Cats

func GetAllCat(c *gin.Context) {
	err := cats.FindAllCat()
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.JSON(200, cats.Cats)
}

func GetCat(c *gin.Context) {
	id := c.Param("id")
	err := cat.FindById(id)
	if err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, cat)

}

func PostCat(c *gin.Context) {
	typeCat := c.PostForm("type")
	models := c.PostForm("model")

	cat, _ = logic.CreatedCat(typeCat, models)
	err := cat.Saver()
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.JSON(202, cat)
	return
}

func DeleteCat(c *gin.Context) {
	id := c.Param("id")
	err := cat.FindById(id)
	if err != nil {
		c.JSON(404, err)
		return
	}
	errDelete := cat.Delete()
	if errDelete != nil {
		c.JSON(503, errDelete)
		return
	}
	c.JSON(200, gin.H{
		"Delete": "ok",
	})
}

func UpdateCat(c *gin.Context) {
	id := c.Param("id")
	typeCat := c.PostForm("type")
	models := c.PostForm("model")

	err := app.FindById(id)
	if err != nil {
		c.JSON(404, err)
		return
	}
	cat.Update(typeCat, models)
	errsaver := app.Saver()
	if errsaver != nil {
		c.JSON(503, errsaver)
		return
	}
	c.JSON(202, cat)

}
