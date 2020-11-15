package Controller

import (
	"github.com/Averdenal/Dotation/Models"
	"github.com/Averdenal/Dotation/logic"
	"github.com/gin-gonic/gin"
)

var app Models.Application

func GetAllApplication(c *gin.Context) {
	var apps Models.Applications
	err := apps.FindAllApplication()
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.JSON(200, apps)
}

func GetApplication(c *gin.Context) {
	id := c.Param("id")

	err := app.FindById(id)
	if err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, app)

}

func PostApplication(c *gin.Context) {
	nom := c.PostForm("name")
	version := c.PostForm("version")

	app, _ = logic.CreatedApplication(nom, version)
	err := app.Saver()
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.JSON(202, app)
	return
}

func DeleteApplication(c *gin.Context) {
	id := c.Param("id")
	err := app.FindById(id)
	if err != nil {
		c.JSON(404, err)
		return
	}
	errDelete := app.Delete()
	if errDelete != nil {
		c.JSON(503, errDelete)
		return
	}
	c.JSON(200, gin.H{
		"Delete": "ok",
	})
}

func UpdateApplication(c *gin.Context) {
	id := c.Param("id")
	nom := c.PostForm("name")
	version := c.PostForm("version")

	err := app.FindById(id)
	if err != nil {
		c.JSON(404, err)
		return
	}
	app.Update(nom, version)
	errsaver := app.Saver()
	if errsaver != nil {
		c.JSON(503, errsaver)
		return
	}
	c.JSON(202, app)

}
