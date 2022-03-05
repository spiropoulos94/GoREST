package controllers

import (
	"encoding/json"
	"go-api/models"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateList(c *gin.Context) {
	var list models.List

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(jsonData, &list)

	list.Name = strings.Trim(list.Name, " ")

	if list.Name == "" {
		c.JSON(400, gin.H{
			"message": "List name is needed",
		})
		return
	}

	// if len(list.Items) == 0 {
	// 	c.JSON(400, gin.H{
	// 		"message": "Cannot Create empty list",
	// 	})
	// 	return
	// }

	result := models.DB.Create(&list)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "List created",
		"data":    list,
	})
}

func FindList(c *gin.Context) {

	id := c.Param("id")

	var list models.List
	result := models.DB.First(&list, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "List found",
		"data":    list,
	})
}

func GetLists(c *gin.Context) {
	var lists []models.List
	result := models.DB.Find(&lists)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "lists successfully retrieved",
		"data":    lists,
	})
}

func DeleteList(c *gin.Context) {
	id := c.Param("id")

	queryResult := models.DB.First(&models.List{}, id)

	if queryResult.Error != nil {
		c.JSON(404, gin.H{
			"message": queryResult.Error.Error(),
		})
		return
	}

	result := models.DB.Unscoped().Delete(&models.List{}, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": queryResult.Error.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "List deleted!",
		})
	}
}
