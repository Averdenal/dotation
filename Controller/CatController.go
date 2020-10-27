package Controller

import (
	"github.com/Averdenal/Dotation/Models"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	var cats Models.Cats
	cats.FindAllCat()
	c.JSON(200, cats)
}
