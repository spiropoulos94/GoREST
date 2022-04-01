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

	api := router.Group("/api", CheckHeaderForJWT())
	{
		userGroup := api.Group("/user")
		{
			// user group handlers
			userGroup.GET("/", controllers.GetUsers)
			userGroup.GET("/:id", controllers.FindUser)
			userGroup.POST("/", controllers.CreateUser)
			userGroup.DELETE("/:id", controllers.DeleteUser)
		}

		listGroup := api.Group("/list")
		{
			// todo group handlers
			listGroup.GET("/", controllers.GetLists)
			listGroup.POST("/", controllers.CreateList)
			listGroup.GET("/:id", controllers.FindList)
			listGroup.DELETE("/:id", controllers.DeleteList)
			listGroup.PUT("/:id", controllers.UpdateList)
		}
	}

	// Auth routes
	router.POST("/signup", controllers.Signup)
	router.POST("/signin", controllers.Signin)

	router.Run()
}
