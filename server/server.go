package server

import (
	"github.com/Averdenal/Dotation/Controller"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

//Server API
func Server() {
	//r.Static("/assets", "./assets")

	//r.GET("/test", Controller.Test)
	ordinateur := r.Group("/ordinateurs")
	{
		ordinateur.GET("", Controller.GetAllOrdinateur)
		ordinateur.GET("/:id", Controller.GetOrdinateur)
		ordinateur.POST("", Controller.PostOrdinateur)
		ordinateur.DELETE("/:id", Controller.DeleteOrdinateur)
		ordinateur.PUT("/:id", Controller.UpdateOrdinateur)
	}

	user := r.Group("/users")
	{
		user.GET("", Controller.GetAllUser)
		user.GET("/:id", Controller.GetUser)
		user.POST("", Controller.PostUser)
		user.DELETE("/:id", Controller.DeleteUser)
		user.PUT("/:id", Controller.UpdateUser)
	}

	materiel := r.Group("/materiels")
	{
		materiel.GET("", Controller.GetAllMateriel)
		materiel.GET("/:id", Controller.GetMateriel)
		materiel.POST("", Controller.PostMateriel)
		materiel.DELETE("/:id", Controller.DeleteMateriel)
		materiel.PUT("/:id", Controller.UpdateMateriel)
	}

	application := r.Group("/applications")
	{
		application.GET("", Controller.GetAllApplication)
		application.GET("/:id", Controller.GetApplication)
		application.POST("", Controller.PostApplication)
		application.DELETE("/:id", Controller.DeleteApplication)
		application.PUT("/:id", Controller.UpdateApplication)
	}
	cat := r.Group("/cats")
	{
		cat.GET("", Controller.GetAllCat)
		cat.GET("/:id", Controller.GetCat)
		cat.POST("", Controller.PostCat)
		cat.DELETE("/:id", Controller.DeleteCat)
		cat.PUT("/:id", Controller.UpdateCat)
	}
	servererr := r.Run()
	if servererr != nil {
		panic(servererr)
	}
}
