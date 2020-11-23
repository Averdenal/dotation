package Controller

import (
	"encoding/json"
	"github.com/Averdenal/Dotation/Models"
	"github.com/gin-gonic/gin"
)

func GetAllOs(c *gin.Context) {
	var oss Models.Opss
	err := oss.FindAllOs()
	if err != nil {
		c.AbortWithStatusJSON(503, err)
	}
	c.JSON(200, oss.Ops)
}
func PostOs(c *gin.Context) {
	var CreadOs Models.CredsOs
	err := json.NewDecoder(c.Request.Body).Decode(&CreadOs)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"err": "information erronée"})
	}
	os := Models.Ops{Version: CreadOs.Version, Name: CreadOs.Name}
	errSaver := os.Saver()
	if errSaver != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"err": "erreur de savegarde",
		})
	}
	c.JSON(202, os)
}
