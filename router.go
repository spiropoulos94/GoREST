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

	user := router.Group("/user")
	list := router.Group("/list")

	// User Routes
	user.GET("/", controllers.GetUsers)
	user.GET("/:id", controllers.FindUser)
	user.POST("/", controllers.CreateUser)
	user.DELETE("/:id", controllers.DeleteUser)
	// List Routes
	list.GET("/", controllers.GetLists)
	list.POST("/", controllers.CreateList)
	list.GET("/:id", controllers.FindList)
	list.DELETE("/:id", controllers.DeleteList)
	list.PUT("/:id", controllers.UpdateList)

	router.Run()
}
