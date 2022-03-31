package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	result := models.DB.Find(&users)
	fmt.Println("result", result.RowsAffected)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error,
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
			"message": result.Error.Error(),
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

	newUser.Name = strings.Trim(newUser.Name, " ")

	if newUser.Name == "" {
		fmt.Println("den exeiu onoma")
		c.JSON(400, gin.H{
			"message": "Please add user name",
		})
	}

	result := models.DB.Create(&newUser)

	if result.Error != nil {
		fmt.Println("error creating user")
		c.JSON(400, gin.H{
			"message": result.Error,
		})
		return
	}

	fmt.Println("User Created, Rows Affected: ", result.RowsAffected)
	c.JSON(200, gin.H{
		"message": "User created successfuly",
		"user":    newUser,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	queryResult := models.DB.First(&models.User{}, id)

	if queryResult.Error != nil {
		c.JSON(404, gin.H{
			"message": queryResult.Error.Error(),
		})
		return
	}

	fmt.Println(queryResult)

	result := models.DB.Unscoped().Delete(&models.User{}, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": queryResult.Error.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "User deleted!",
		})
	}
}
