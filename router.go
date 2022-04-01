package main

import (
	"go-api/controllers"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func routerStart() {
	router := gin.Default()

	models.SetupDatabase()
	models.MakeTables(false)

	//groupare sto "/api" olo to user kai to list

	api := router.Group("/api")
	{
		api.GET("user/", controllers.GetUsers)
		api.GET("user/:id", controllers.FindUser)
		api.POST("user/", controllers.CreateUser)
		api.DELETE("user/:id", controllers.DeleteUser)

		api.GET("list/", controllers.GetLists)
		api.POST("list/", controllers.CreateList)
		api.GET("list/:id", controllers.FindList)
		api.DELETE("list/:id", controllers.DeleteList)
		api.PUT("list/:id", controllers.UpdateList)
	}

	// // User Routes
	// user.GET("/", controllers.GetUsers)
	// user.GET("/:id", controllers.FindUser)
	// user.POST("/", controllers.CreateUser)
	// user.DELETE("/:id", controllers.DeleteUser)
	// // List Routes
	// list.GET("/", controllers.GetLists)
	// list.POST("/", controllers.CreateList)
	// list.GET("/:id", controllers.FindList)
	// list.DELETE("/:id", controllers.DeleteList)
	// list.PUT("/:id", controllers.UpdateList)

	// Auth routes
	router.POST("/signup", controllers.Signup)
	router.POST("/signin", controllers.Signin)

	router.Run()
}
