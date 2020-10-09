package server

import (
	"fmt"

	"github.com/Averdenal/Dotation/Controller"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Server() {
	v1 := r.Group("/v1")
	{
		v1.GET("/", DemoFunc)
		v1.POST("/Dotation", Controller.NewDotation)
	}
	r.Run()
}
func DemoFunc(c *gin.Context) {
	fmt.Println("demo")
}
