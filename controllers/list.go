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
			"message": "could not create list",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "List created",
		"data":    list,
	})
}

func GetList(c *gin.Context) {

	id := c.Param("id")

	var list models.List
	result := models.DB.First(&list, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Could not find list",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "List found",
		"data":    list,
	})
}
