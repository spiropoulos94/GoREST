package main

// Kane refactor se models kai conrollers opws edw : https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/

import (
	"go-api/controllers"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	models.SetupDatabase()
	models.MakeTables(false)

	// User Routes
	r.GET("/user", controllers.GetUsers)
	r.GET("/user/:id", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	// List Routes
	r.GET("/list", controllers.GetLists)
	r.POST("/list", controllers.CreateList)
	r.GET("/list/:id", controllers.FindList)
	r.DELETE("/list/:id", controllers.DeleteList)
	r.PUT("/list/:id", controllers.UpdateList)

	r.Run()
}
