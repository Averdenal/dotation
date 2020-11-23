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
	return
}
func GetOs(c *gin.Context) {
	var os Models.Ops
	id := c.Param("id")
	errFind := os.FindById(id)
	if errFind != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"err": "not found",
		})
	}
	c.JSON(200, os)
	return

}
func PostOs(c *gin.Context) {
	var CredsOs Models.CredsOs
	err := json.NewDecoder(c.Request.Body).Decode(&CredsOs)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"err": "information erronée"})
	}
	os := Models.Ops{Version: CredsOs.Version, Name: CredsOs.Name}
	errSaver := os.Saver()
	if errSaver != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"err": "erreur de savegarde",
		})
	}
	c.JSON(202, os)
	return
}
func UpdateOs(c *gin.Context) {
	var os Models.Ops
	id := c.Param("id")
	errFind := os.FindById(id)
	if errFind != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"err": "not found",
		})
	}
	var CredsOs Models.CredsOs
	err := json.NewDecoder(c.Request.Body).Decode(&CredsOs)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"err": "information erronée"})
	}
	os.Name = CredsOs.Name
	os.Version = CredsOs.Version
	errSaver := os.Saver()
	if errSaver != nil {
		c.AbortWithStatusJSON(503, errSaver)
	}
	c.JSON(200, os)
	return
}

func DeleteOs(c *gin.Context) {
	var os Models.Ops
	id := c.Param("id")
	errFind := os.FindById(id)
	if errFind != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"err": "not found",
		})
	}
	errDelete := os.Delete()
	if errDelete != nil {
		c.AbortWithStatusJSON(503, errDelete)
	}
	c.JSON(202, "ok")
}
