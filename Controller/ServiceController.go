package Controller

import (
	"fmt"
	"github.com/Averdenal/Dotation/Models"
	"github.com/gin-gonic/gin"
)

func GetAllService(c *gin.Context) {
	fmt.Println(c)
	s := Models.Services{}
	err := s.FindAll()
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.JSON(200, s)
}

func GetService(c *gin.Context) {

}

func PostService(c *gin.Context) {

}

func DeleteService(c *gin.Context) {

}

func UpdateService(c *gin.Context) {

}
