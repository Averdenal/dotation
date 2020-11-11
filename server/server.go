package server

import (
	"github.com/Averdenal/Dotation/Controller"
	"github.com/Averdenal/Dotation/logic"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

//Server API
func Server() {

	security := r.Group("")
	{
		security.POST("/register", CORS(), Controller.Register)
		security.POST("/login", CORS(), Controller.Login)
		security.GET("/Welcome", CORS(), Controller.Welcome)
	}

	service := r.Group("/services")
	service.Use(CORS(), logic.Access())
	{
		service.GET("", Controller.GetAllService)
		//service.GET("/:id", Controller.GetOrdinateur)
		//service.POST("", Controller.PostOrdinateur)
		//service.DELETE("/:id", Controller.DeleteOrdinateur)
		//service.PUT("/:id", Controller.UpdateOrdinateur)
	}

	ordinateur := r.Group("/ordinateurs ")
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

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true

	r.Use(cors.New(config))

	servererr := r.Run()
	if servererr != nil {
		panic(servererr)
	}
}
