package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	result := models.DB.Find(&users)
	fmt.Println("result", result.RowsAffected)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "could not retrieve data",
		})
	}

	c.JSON(200, gin.H{
		"message": "Users successfully retrieved",
		"data":    users,
	})
}

func FindUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User
	result := models.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Could not find user",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User found",
		"data":    user,
	})
}

func CreateUser(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	newUser := models.User{}
	json.Unmarshal(jsonData, &newUser)

	result := models.DB.Create(&newUser)

	if result.Error != nil {
		fmt.Println("error creating user")
		c.JSON(400, gin.H{
			"message": "Error Creating new record",
		})
	}

	fmt.Println("User Created, Rows Affected: ", result.RowsAffected)
	c.JSON(200, gin.H{
		"message": "User created successfuly",
		"user":    newUser,
	})
}
