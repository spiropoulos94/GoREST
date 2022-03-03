package main

// Kane refactor se models kai conrollers opws edw : https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/

import (
	"go-api/controllers"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	models.SetupDatabase()
	models.MakeTables(false)

	router.GET("/user", controllers.GetUsers)
	router.GET("/user/:id", controllers.FindUser)
	router.POST("/user", controllers.CreateUser)

	router.Run()
}
