package server

import (
	"github.com/Averdenal/Dotation/Controller"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

//Server API
func Server() {
	r.Static("/assets", "./assets")

	r.GET("/", Controller.Home)

	ordinateur := r.Group("/ordinateur")
	{
		ordinateur.GET("", Controller.GetAllOrdinateur)
		ordinateur.GET("/:id", Controller.GetByOrdinateur)
		ordinateur.POST("", Controller.PostOrdinateur)
		ordinateur.DELETE("/:id", Controller.DeleteOrdinateur)
		ordinateur.PUT("/:id", Controller.UpdateOrdinateur)
	}

	user := r.Group("/user")
	{
		user.GET("", Controller.NewUser)
		user.GET("/:id", Controller.NewUser)
		user.POST("", Controller.NewUser)
		user.DELETE("/:id", Controller.NewUser)
		user.PUT("/:id", Controller.NewUser)
	}

	materiel := r.Group("/materiels")
	{
		materiel.GET("", Controller.NewUser)
		materiel.GET("/:id", Controller.NewUser)
		materiel.POST("", Controller.NewUser)
		materiel.DELETE("/:id", Controller.NewUser)
		materiel.PUT("/:id", Controller.NewUser)
	}

	application := r.Group("/materiels")
	{
		application.GET("", Controller.NewUser)
		application.GET("/:id", Controller.NewUser)
		application.POST("", Controller.NewUser)
		application.DELETE("/:id", Controller.NewUser)
		application.PUT("/:id", Controller.NewUser)
	}
	r.Run()
}
